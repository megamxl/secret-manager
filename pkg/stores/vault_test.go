package stores

import (
	"context"
	"testing"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema" // Add this import
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	vaultContainer "github.com/testcontainers/testcontainers-go/modules/vault"
)

func TestVaultBackend_Integration(t *testing.T) {
	ctx := context.Background()
	rootToken := "root-token"

	// 1. Start Vault Container
	vaultC, err := vaultContainer.Run(ctx, "hashicorp/vault:1.13",
		testcontainers.WithEnv(map[string]string{
			"VAULT_DEV_ROOT_TOKEN_ID": rootToken,
		}),
	)
	require.NoError(t, err)
	defer func() { _ = vaultC.Terminate(ctx) }()

	endpoint, err := vaultC.HttpHostAddress(ctx)
	require.NoError(t, err)

	// 2. Setup Backend
	backend := &VaultBackend{}
	config := &Config{
		Type: "hc-vault",
		URL:  endpoint,
		Auth: map[string]string{"token": rootToken},
	}
	err = backend.InstateClient(config)
	require.NoError(t, err)

	// 3. Prepare Test Data (FIXED TYPING)
	secretPath := "mysql"
	mountPath := "secret"

	_, err = backend.client.Secrets.KvV2Write(ctx, secretPath,
		schema.KvV2WriteRequest{
			Data: map[string]interface{}{
				"password": "super-secret-password",
				"user":     "admin",
			},
		},
		vault.WithMountPath(mountPath),
	)
	require.NoError(t, err)

	// 4. Test ResolveSecrets logic
	t.Run("Successfully resolve mapped secrets", func(t *testing.T) {
		parsedSecret := ParsedSecret{
			MountPath: mountPath,
			Refs: []SecretRef{
				{
					TemplateName:                "DB_PASS",
					RemotePathWithoutSecretName: secretPath,
					RemoteSecretName:            "password",
				},
			},
		}

		results, err := backend.ResolveSecrets(parsedSecret)
		assert.NoError(t, err)
		assert.Equal(t, "super-secret-password", results["DB_PASS"])
	})
}
