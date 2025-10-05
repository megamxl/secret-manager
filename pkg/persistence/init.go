package persistence

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	var err error

	// Use modernc.org/sqlite instead of cgo sqlite3
	db, err := gorm.Open(sqlite.Open("file:store_configs.db?cache=shared&mode=rwc"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&StoreConfig{}, &ManagedFiles{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
