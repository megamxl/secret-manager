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
	ID       uint      `gorm:"primaryKey"`
	Filepath string    `gorm:"uniqueIndex"`
	Request  string    `gorm:"type:TEXT"`
	Reroll   time.Time `gorm:"index"`
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
		Filepath: config.FilePath,
		Request:  string(jsonData),
		Reroll:   time.Now().Add(time.Second * 10),
	}

	// Use OnConflict to handle updates on unique key violations
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "filepath"}},                     // The column with the unique index
		DoUpdates: clause.AssignmentColumns([]string{"request", "reroll"}), // Columns to update
	}).Create(&managedFile).Error
}

func GetPendingRerolls(db *gorm.DB) ([]types.CreateSecretRequest, error) {
	var records []ManagedFiles
	var results []types.CreateSecretRequest

	// We look for records where the Reroll time is in the future
	// (Current Time < Reroll)
	err := db.Where("reroll <= ?", time.Now()).Find(&records).Error

	if err != nil {
		return nil, err
	}

	for _, rec := range records {
		var req types.CreateSecretRequest
		// Unmarshal the JSON string back into the struct
		if err := json.Unmarshal([]byte(rec.Request), &req); err != nil {
			log.Println("Error unmarshalling request for reroll:", err)
			continue
		}
		results = append(results, req)
	}

	return results, nil
}
