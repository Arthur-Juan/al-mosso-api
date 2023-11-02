package config

import (
	"al-mosso-api/pkg/database/schemas"
	"al-mosso-api/pkg/logger"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePg() (*gorm.DB, error) {
	logger := logger.GetLogger("db")
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("ERROR TO START DB: %s", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Appointment{}, &schemas.Chef{}, &schemas.Client{}, &schemas.Food{}, schemas.Contact{},
		&schemas.Config{})

	if err != nil {
		logger.Errorf("Migrating ERROR: %v", err)
		return nil, err
	}

	logger.Info("Migration made!")
	return db, nil
}
