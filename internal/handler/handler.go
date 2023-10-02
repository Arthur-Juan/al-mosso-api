package handler

import (
	"al-mosso-api/config"
	logger2 "al-mosso-api/pkg/logger"
	"gorm.io/gorm"
)

var (
	logger *logger2.Logger
	db     *gorm.DB
)

func InitHandlerConfig() {
	logger = logger2.GetLogger("handler")
	db = config.GetDb()
}
