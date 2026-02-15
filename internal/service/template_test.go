package service

import (
	"errors"
	"secret-manager/pkg/persistence"
	"secret-manager/pkg/stores"
	"secret-manager/pkg/types"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type MockBackend struct {
	ResolvedValues map[string]string
	ShouldFail     bool
}

func (m *MockBackend) InstateClient(config *stores.Config) error { return nil }
func (m *MockBackend) ResolveSecrets(s stores.ParsedSecret) (map[string]string, error) {
	if m.ShouldFail {
		return nil, errors.New("mock backend failure")
	}
	return m.ResolvedValues, nil
}

func TestToParsedSecret(t *testing.T) {
	request := types.CreateSecretRequest{
		FilePath:  "/etc/config.conf",
		StoreName: "vault-prod",
		SecretWrapper: types.Secret{
			Content: "pass={{.db_pass}}",
			SecretReferences: map[string]string{
				"db_pass": "kv/data/mysql/root",
			},
		},
	}
	config := stores.Config{MountPath: "secret/"}

	parsed := ToParsedSecret(request, config)

	assert.Equal(t, "/etc/config.conf", parsed.FilePath)
	assert.Equal(t, "secret/", parsed.MountPath)
	require.Len(t, parsed.Refs, 1)

	// Testing the split logic inside generateSecretRefs
	assert.Equal(t, "db_pass", parsed.Refs[0].TemplateName)
	assert.Equal(t, "kv/data/mysql", parsed.Refs[0].RemotePathWithoutSecretName)
	assert.Equal(t, "root", parsed.Refs[0].RemoteSecretName)
}

func TestTemplateServiceImpl_BasicOperations(t *testing.T) {
	// Setup stable in-memory DB
	dsn := "file:template_test?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	require.NoError(t, err)
	_ = db.AutoMigrate(&persistence.ManagedFiles{}, &persistence.StoreConfig{})

	svc := TemplateServiceImpl{Db: db}

	t.Run("Store and Delete Config", func(t *testing.T) {
		req := types.CreateSecretRequest{Name: "test-svc", FilePath: "/tmp/test"}

		// Test Save
		err := svc.StoreSecretConfig(req)
		assert.NoError(t, err)

		// Test Delete
		err = svc.DeleteSecretConfig("test-svc")
		assert.NoError(t, err)

		// Verify deletion
		_, err = persistence.FindConfigByName(db, "test-svc")
		assert.Error(t, err)
	})
}
