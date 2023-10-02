package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
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

func GetDb() *gorm.DB {
	return db
}
