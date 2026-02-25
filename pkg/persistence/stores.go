package persistence

import (
	"encoding/json"
	"errors"
	"secret-manager/internal/logging"
	"secret-manager/pkg/stores"

	"gorm.io/gorm"
)

type StoreConfig struct {
	Name   string `gorm:"primaryKey"`
	Config string `gorm:"type:TEXT"`
}

func SaveStoreConfig(db *gorm.DB, config stores.Config) error {

	if db == nil {
		return errors.New("database not initialized")
	}
	jsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}
	return db.Save(&StoreConfig{
		Name:   config.ReferenceName,
		Config: string(jsonData),
	}).Error

}

func GetAllStores(db *gorm.DB) ([]stores.Config, error) {

	if db == nil {
		return nil, errors.New("database not initialized")
	}
	var dbConfigs []StoreConfig
	if err := db.Find(&dbConfigs).Error; err != nil {
		return nil, err
	}

	results := make([]stores.Config, 0, len(dbConfigs))

	for _, dbEntry := range dbConfigs {
		var c stores.Config
		if err := json.Unmarshal([]byte(dbEntry.Config), &c); err != nil {
			logging.L.App.Warn("Error unmarshalling store configuration")
			continue
		}
		results = append(results, c)
	}

	return results, nil
}

func FindConfigByName(db *gorm.DB, name string) (stores.Config, error) {

	if db == nil {
		return stores.Config{}, errors.New("database not initialized")
	}
	var store StoreConfig
	if err := db.First(&store, "name = ?", name).Error; err != nil {
		return stores.Config{}, err
	}
	var config stores.Config
	err := json.Unmarshal([]byte(store.Config), &config)
	return config, err
}

func DeleteConfigByName(db *gorm.DB, name string) error {
	if db == nil {
		return errors.New("database not initialized")
	}
	return db.Delete(&StoreConfig{}, "name = ?", name).Error
}

func UpdateStoreConfig(db *gorm.DB, config stores.Config) error {
	if db == nil {
		return errors.New("database not initialized")
	}

	jsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	result := db.Model(&StoreConfig{}).
		Where("name = ?", config.ReferenceName).
		Updates(map[string]interface{}{
			"config": string(jsonData),
		})

	if result.Error != nil {
		return result.Error
	}

	// If no rows were affected, the store didn't exist
	if result.RowsAffected == 0 {
		return errors.New("store config not found")
	}

	return nil
}
