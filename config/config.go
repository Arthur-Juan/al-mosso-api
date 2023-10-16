package config

import (
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	filePath = "./uploads"
)

func Init() error {
	var err error
	//init sqlite
	db, err = InitializeSqlite()

	if err != nil {
		return err
	}

	return nil
}

func GetFilePath() string {
	return filePath
}

func GetDb() *gorm.DB {
	return db
}

func GetHostName() string {
	return "http://localhost:8080"
}
