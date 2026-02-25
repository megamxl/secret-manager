package persistence

import (
	"encoding/json"
	"errors"
	"fmt"
	"secret-manager/internal/logging"
	"secret-manager/pkg/types"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ManagedFiles represents a request with its associated paths
type ManagedFiles struct {
	ID      uint      `gorm:"primaryKey"`
	Name    string    `gorm:"uniqueIndex"`
	Request string    `gorm:"type:TEXT"`
	Reroll  time.Time `gorm:"index"`
}

func SaveSecretConfig(db *gorm.DB, config types.CreateSecretRequest) error {
	if db == nil {
		return errors.New("database not initialized")
	}

	jsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	managedFile := ManagedFiles{
		Name:    config.Name,
		Request: string(jsonData),
		Reroll:  time.Now().Add(time.Duration(config.RerollTime)),
	}

	// Use OnConflict to handle updates on unique key violations
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"request", "reroll"}),
	}).Create(&managedFile).Error
}

func GetRequestByName(db *gorm.DB, name string) (*types.CreateSecretRequest, error) {
	var record ManagedFiles
	var result types.CreateSecretRequest

	// Lookup the record by the name column
	err := db.Where("name = ?", name).First(&record).Error
	if err != nil {
		return nil, err
	}

	// Unmarshal the stored JSON string into your response struct
	err = json.Unmarshal([]byte(record.Request), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal request for %s: %w", name, err)
	}

	return &result, nil
}

func GetPendingRerolls(db *gorm.DB) ([]types.CreateSecretRequest, error) {
	var records []ManagedFiles
	var results []types.CreateSecretRequest

	err := db.Where("reroll <= ?", time.Now()).Find(&records).Error

	if err != nil {
		return nil, err
	}

	for _, rec := range records {
		var req types.CreateSecretRequest
		// Unmarshal the JSON string back into the struct only json since we only persist json
		if err := json.Unmarshal([]byte(rec.Request), &req); err != nil {
			logging.L.App.Error(fmt.Sprintf("Error unmarshalling request for reroll: %s", err))
			continue
		}
		results = append(results, req)
	}

	return results, nil
}

func UpdateSecretConfig(db *gorm.DB, config types.CreateSecretRequest) error {
	if db == nil {
		return errors.New("database not initialized")
	}

	jsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	result := db.Model(&ManagedFiles{}).
		Where("name = ?", config.Name).
		Updates(map[string]interface{}{
			"request": string(jsonData),
			//Setting re-roll-time to now that the updated secret gets rerolled and used
			"reroll": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no record found with that name, please create that file")
	}

	return nil
}

func DeleteSecretConfig(db *gorm.DB, name string) error {
	if db == nil {
		return errors.New("database not initialized")
	}

	result := db.Where("name = ?", name).Delete(&ManagedFiles{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no record found with that name to delete")
	}

	return nil
}

func IsStoreInUse(db *gorm.DB, storeName string) (bool, error) {
	var count int64

	// It looks for the key "store_name" inside the "request" column
	err := db.Model(&ManagedFiles{}).
		Where("request ->> '$.store_name' = ?", storeName).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
