package service

import (
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"secret-manager/pkg/persistence"
	"secret-manager/pkg/types"
	"testing"
	"time"

	"errors"
)

// MockTemplateService satisfies the TemplateService interface
type MockTemplateService struct {
	OnFetch func(req types.CreateSecretRequest) error
}

func (m MockTemplateService) FetchAndStoreTemplate(req types.CreateSecretRequest) error {
	if m.OnFetch != nil {
		return m.OnFetch(req)
	}
	return nil
}

func SetupRotationDB(t *testing.T) *gorm.DB {
	dsn := "file:memdb1?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Ensure we migrate BOTH models used in the persistence layer
	err = db.AutoMigrate(&persistence.ManagedFiles{}, &persistence.StoreConfig{})
	require.NoError(t, err)

	return db
}

func TestRotationJob_Run(t *testing.T) {
	db := SetupRotationDB(t)

	t.Run("Successfully processes all pending rerolls", func(t *testing.T) {
		// Prepare data: 3 secrets that are "expired" (reroll in the past)
		secrets := []types.CreateSecretRequest{
			{Name: "sec1", FilePath: "/tmp/1", RerollTime: types.RerollDuration(time.Hour)},
			{Name: "sec2", FilePath: "/tmp/2", RerollTime: types.RerollDuration(time.Hour)},
		}

		for _, s := range secrets {
			err := persistence.SaveSecretConfig(db, s)
			require.NoError(t, err)
		}

		// Force them to be "pending" by moving the reroll date to the past
		db.Model(&persistence.ManagedFiles{}).Where("1=1").Update("reroll", time.Now().Add(-1*time.Minute))

		// Setup Mock
		callCount := 0
		mockSvc := MockTemplateService{
			OnFetch: func(req types.CreateSecretRequest) error {
				callCount++
				return nil
			},
		}

		job := RotationJob{Db: db, Service: mockSvc}
		job.Run()

		assert.Equal(t, 2, callCount, "Both secrets should have been processed")

		// Verify reroll time was updated in DB (no longer pending)
		pending, _ := persistence.GetPendingRerolls(db)
		assert.Len(t, pending, 0, "No secrets should be pending after Run()")
	})

	t.Run("Handles partial failures correctly", func(t *testing.T) {
		// Insert one secret
		s := types.CreateSecretRequest{Name: "fail-test", FilePath: "/tmp/fail"}
		persistence.SaveSecretConfig(db, s)
		db.Model(&persistence.ManagedFiles{}).Update("reroll", time.Now().Add(-1*time.Minute))

		// Mock returns an error
		mockSvc := MockTemplateService{
			OnFetch: func(req types.CreateSecretRequest) error {
				return errors.New("network timeout")
			},
		}

		job := RotationJob{Db: db, Service: mockSvc}
		job.Run()

		// The job logs the error but doesn't crash.
		// Since it failed, SaveSecretConfig wasn't called to update the reroll time,
		// so it should STILL be pending.
		pending, _ := persistence.GetPendingRerolls(db)
		assert.Len(t, pending, 1, "Failed reroll should remain pending for next retry")
	})
}
