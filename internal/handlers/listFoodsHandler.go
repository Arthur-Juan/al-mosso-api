package handlers

import (
	"al-mosso-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

func ListFoodHandler(ctx *fiber.Ctx) error {

	svc := services.NewListFoodsService()
	result, err := svc.Exceute()

	if err != nil {
		return DispatchError(ctx, *err)
	}

	if result == nil {
		return NotFound(ctx, "No foods found")
	}

	return Ok(ctx, result)

}
