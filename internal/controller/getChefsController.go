package controller

import (
	"al-mosso-api/internal/handler"
	logger2 "al-mosso-api/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func GetChefsController(ctx *fiber.Ctx) error {
	result, err := handler.GetChefsHandler()
	logger := logger2.NewLogger("controller")
	logger.Infof("controller: %v", result)
	if err != nil {
		return InternalServerError(ctx, err)
	}
	if result == nil {
		return NotFound(ctx, "no chefs found")
	}

	return Ok(ctx, result)
}
