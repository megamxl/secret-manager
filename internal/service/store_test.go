package service

import (
	"secret-manager/pkg/persistence"
	"secret-manager/pkg/stores"
	"secret-manager/pkg/types"
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
	err = db.AutoMigrate(&persistence.StoreConfig{}, &persistence.ManagedFiles{})
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

func TestDeleteStoreInUse(t *testing.T) {
	db := SetupStoreTestDB(t)
	service := StoreService{Db: db}

	t.Run("Should fail to delete store when used in ManagedFiles JSON", func(t *testing.T) {
		storeName := "production-vault"

		// 1. Create a store
		err := service.CreateStore(stores.Config{ReferenceName: storeName})
		require.NoError(t, err)

		// 2. Create a ManagedFile that "uses" this store inside the JSON blob
		// Note: Make sure SaveSecretConfig puts storeName into the 'Request' field
		req := types.CreateSecretRequest{
			Name:      "api-key",
			StoreName: storeName,
		}
		err = persistence.SaveSecretConfig(db, req)
		require.NoError(t, err)

		// 3. Attempt to delete through the service
		err = service.DeleteStore(storeName)

		// 4. Assertions
		assert.Error(t, err, "Service should return an error because the store is in use")
		assert.Contains(t, err.Error(), "still in use", "Error message should mention the store is in use")

		// 5. Double check the store still exists in DB
		_, findErr := persistence.FindConfigByName(db, storeName)
		assert.NoError(t, findErr, "Store should NOT have been deleted from the database")
	})

	t.Run("Should succeed to delete store when NOT in use", func(t *testing.T) {
		storeName := "unused-vault"
		service.CreateStore(stores.Config{ReferenceName: storeName})

		err := service.DeleteStore(storeName)
		assert.NoError(t, err, "Should delete successfully when no managed files reference it")
	})
}
