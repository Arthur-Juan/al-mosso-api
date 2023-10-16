package config

import (
	"al-mosso-api/pkg/database/schemas"
	logger2 "al-mosso-api/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := logger2.GetLogger("sqlite")

	_, err := os.Stat("./db/main.db")
	if os.IsNotExist(err) {
		logger.Info("Database not found, migrating then")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create("./db/main.db")
		if err != nil {
			return nil, err
		}
		err = file.Close()
		if err != nil {
			return nil, err
		}

	}

	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})

	if err != nil {
		logger.Errorf("Init sqlite ERROR: %v", err)
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
