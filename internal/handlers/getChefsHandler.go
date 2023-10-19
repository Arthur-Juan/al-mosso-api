package handlers

import (
	"al-mosso-api/internal/services"
	logger2 "al-mosso-api/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func GetChefsHandler(ctx *fiber.Ctx) error {
	svc := services.NewGetChefsService()
	result, err := svc.Execute()
	logger := logger2.NewLogger("handlers")
	logger.Infof("handlers: %v", result)
	if err != nil {
		return InternalServerError(ctx, err)
	}
	if result == nil {
		return NotFound(ctx, "no chefs found")
	}

	return Ok(ctx, result)
}
