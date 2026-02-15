package service

import (
	"secret-manager/pkg/persistence"
	"secret-manager/pkg/stores"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func SetupStoreTestDB(t *testing.T) *gorm.DB {
	// Use shared memory to prevent the table from disappearing between service calls
	dsn := "file:store_test?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Crucial: Migrate the exact schema used by the persistence layer
	err = db.AutoMigrate(&persistence.StoreConfig{})
	require.NoError(t, err)

	return db
}

func TestStoreService(t *testing.T) {
	db := SetupStoreTestDB(t)
	svc := StoreService{Db: db}

	t.Run("Create and Retrieve Store", func(t *testing.T) {
		conf := stores.Config{
			ReferenceName: "vault-production",
			// Add other fields based on your stores.Config struct
		}

		// Test Create
		err := svc.CreateStore(conf)
		assert.NoError(t, err)

		// Verify via persistence layer directly to ensure isolation
		retrieved, err := persistence.FindConfigByName(db, "vault-production")
		assert.NoError(t, err)
		assert.Equal(t, "vault-production", retrieved.ReferenceName)
	})

	t.Run("Update Existing Store", func(t *testing.T) {
		conf := stores.Config{ReferenceName: "vault-production"}

		// In a real scenario, you'd change a field here
		err := svc.UpdateStore(conf)
		assert.NoError(t, err)
	})

	t.Run("Update Non-Existent Store Returns Error", func(t *testing.T) {
		conf := stores.Config{ReferenceName: "ghost-vault"}
		err := svc.UpdateStore(conf)

		// This should fail because persistence.UpdateStoreConfig
		// returns an error if RowsAffected == 0
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "store config not found")
	})

	t.Run("Delete Store", func(t *testing.T) {
		err := svc.DeleteStore("vault-production")
		assert.NoError(t, err)

		// Verify it's actually gone
		_, err = persistence.FindConfigByName(db, "vault-production")
		assert.Error(t, err, "Should not find store after deletion")
	})
}
