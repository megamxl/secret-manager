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

	//TODO build and connect Backend before storing

	err := persistence.SaveStoreConfig(s.Db, req)
	if err != nil {
		log.Print("Error creating store:", err)
		return err
	}

	return nil
}

func (s StoreService) UpdateStore(req stores.Config) error {
	err := persistence.UpdateStoreConfig(s.Db, req)
	if err != nil {
		log.Printf("Error updating store %s: %v", req.ReferenceName, err)
		return err
	}
	return nil
}

// DeleteStore removes a store configuration from the database
func (s StoreService) DeleteStore(name string) error {
	err := persistence.DeleteConfigByName(s.Db, name)
	if err != nil {
		log.Printf("Error deleting store %s: %v", name, err)
		return err
	}
	return nil
}
