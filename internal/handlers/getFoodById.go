package handlers

import (
	"al-mosso-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetFoodById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	uId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return InternalServerError(ctx, err)
	}

	svc := services.NewGetFoodByIdService()
	result, err := svc.Execute(uId)
	if err != nil {
		return InternalServerError(ctx, err)
	}
	return Ok(ctx, result)
}
