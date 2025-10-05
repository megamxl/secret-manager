package service

import (
	"log"
	"secret-manager/pkg/persistence"
	"secret-manager/pkg/stores"

	"gorm.io/gorm"
)

type StoreService struct {
	Db *gorm.DB
}

func (s StoreService) CreateStore(req stores.Config) error {

	err := persistence.SaveStoreConfig(s.Db, req)
	if err != nil {
		log.Print("Error creating store:", err)
		return err
	}

	return nil
}
