package controller

import (
	"al-mosso-api/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func ListFoodController(ctx *fiber.Ctx) error {

	result, err := handler.ListFoodsHandler()

	if err != nil {
		return InternalServerError(ctx, err)
	}

	if result == nil {
		return NotFound(ctx, "No foods found")
	}

	return Ok(ctx, result)

}
