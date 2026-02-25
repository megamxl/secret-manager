package persistence

import (
	"secret-manager/internal/config"
	"secret-manager/internal/logging"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	var err error

	newLogger := &logging.GormLogger{
		ZapLogger: logging.L.App,
		LogLevel:  logger.Error,
	}

	// Use modernc.org/sqlite instead of cgo sqlite3
	db, err := gorm.Open(sqlite.Open("file:"+config.Get().DB.Path+"?cache=shared&mode=rwc"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&StoreConfig{}, &ManagedFiles{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
