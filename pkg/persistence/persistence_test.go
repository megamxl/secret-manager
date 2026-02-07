package persistence

import (
	"secret-manager/pkg/stores"
	"secret-manager/pkg/types"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)

	err = db.AutoMigrate(&StoreConfig{}, &ManagedFiles{})
	require.NoError(t, err)

	return db
}

func TestSecretConfigPersistence(t *testing.T) {
	db := SetupTestDB(t)

	t.Run("Full Lifecycle: Save, Reroll Check, and Update", func(t *testing.T) {
		// 1. Save a config with a 2-hour reroll
		initialReq := types.CreateSecretRequest{
			Name:       "api-key",
			StoreName:  "vault-prod",
			RerollTime: types.RerollDuration(2 * time.Hour),
		}

		err := SaveSecretConfig(db, initialReq)
		assert.NoError(t, err)

		// 2. Verify it is NOT pending (since 2 hours is in the future)
		pending, err := GetPendingRerolls(db)
		assert.NoError(t, err)
		assert.Empty(t, pending)

		// 3. Manually force the record into the past to test GetPendingRerolls
		db.Model(&ManagedFiles{}).
			Where("name = ?", "api-key").
			Update("reroll", time.Now().Add(-1*time.Minute))

		pending, err = GetPendingRerolls(db)
		assert.NoError(t, err)
		require.Len(t, pending, 1)
		assert.Equal(t, "api-key", pending[0].Name)
		// Verify custom type unmarshaled correctly from DB nanoseconds
		assert.Equal(t, initialReq.RerollTime, pending[0].RerollTime)

		// 4. Test Update (Should reset reroll to "Now")
		updateReq := initialReq
		updateReq.StoreName = "vault-backup"
		err = UpdateSecretConfig(db, updateReq)
		assert.NoError(t, err)

		// Verify the update took effect
		var record ManagedFiles
		db.First(&record, "name = ?", "api-key")
		assert.Contains(t, record.Request, "vault-backup")
		// Reroll should be ~now, so it should show up in pending again
		pendingNow, _ := GetPendingRerolls(db)
		assert.Len(t, pendingNow, 1)
	})

	t.Run("Delete Non-Existent", func(t *testing.T) {
		err := DeleteSecretConfig(db, "ghost-secret")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no record found")
	})
}

func TestStoreConfigPersistence(t *testing.T) {
	db := SetupTestDB(t)

	t.Run("Save and Update Store", func(t *testing.T) {
		conf := stores.Config{ReferenceName: "primary-vault"}

		// Save
		err := SaveStoreConfig(db, conf)
		assert.NoError(t, err)

		// Update
		err = UpdateStoreConfig(db, stores.Config{ReferenceName: "primary-vault"})
		assert.NoError(t, err)

		// Find
		retrieved, err := FindConfigByName(db, "primary-vault")
		assert.NoError(t, err)
		assert.Equal(t, "primary-vault", retrieved.ReferenceName)

		// Delete
		err = DeleteConfigByName(db, "primary-vault")
		assert.NoError(t, err)
	})
}
