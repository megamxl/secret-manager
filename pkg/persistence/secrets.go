package persistence

import (
	"encoding/json"
	"errors"
	"log"
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
			log.Println("Error unmarshalling request for reroll:", err)
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
