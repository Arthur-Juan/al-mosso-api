package main

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/router"
	"al-mosso-api/internal/services"
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
		return
	}
	services.InitHandlerConfig()
	router.Initialize()
}
