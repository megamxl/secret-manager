package persistence

import (
	"encoding/json"
	"errors"
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
