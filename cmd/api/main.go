package main

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/router"
	"al-mosso-api/internal/services"
	"al-mosso-api/pkg/database"
	logger2 "al-mosso-api/pkg/logger"
)

var (
	logger *logger2.Logger
)

func main() {
	logger = logger2.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization ERROR: %v", err)
		panic(err)
	}
	services.InitHandlerConfig()

	var qtd int

	db := config.GetDb()
	db.Raw("select count(*) from foods").Scan(&qtd)
	logger.Debug(qtd)
	if qtd == 0 {
		database.Seed()
	}
	router.Initialize()
}
