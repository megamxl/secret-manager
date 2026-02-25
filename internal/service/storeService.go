package service

import (
	"errors"
	"fmt"
	"secret-manager/internal/logging"
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
		logging.L.App.Error(fmt.Sprintf("Error creating store: %s", err))
		return err
	}

	return nil
}

func (s StoreService) UpdateStore(req stores.Config) error {
	err := persistence.UpdateStoreConfig(s.Db, req)
	if err != nil {
		logging.L.App.Error(fmt.Sprintf("Error updating store %s: %v", req.ReferenceName, err))
		return err
	}
	return nil
}

// DeleteStore removes a store configuration from the database
func (s StoreService) DeleteStore(name string) error {

	use, err := persistence.IsStoreInUse(s.Db, name)
	if err != nil {
		logging.L.App.Error(fmt.Sprintf("Error Finding store %s: %v", name, err))
		return err
	}

	if use {
		return errors.New(fmt.Sprintf("Error store %s is still in use, remove all secrets using it: %v", name, err))
	}

	err = persistence.DeleteConfigByName(s.Db, name)
	if err != nil {
		logging.L.App.Error(fmt.Sprintf("Error deleting store %s: %v", name, err))
		return err
	}
	return nil
}
