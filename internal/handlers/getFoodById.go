package handlers

import (
	"al-mosso-api/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetFoodById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	uId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return InternalServerError(ctx, err)
	}

	svc := services.NewGetFoodByIdService()
	result, terr := svc.Execute(uId)
	if err != nil {
		return DispatchError(ctx, *terr)
	}
	return Ok(ctx, result)
}
