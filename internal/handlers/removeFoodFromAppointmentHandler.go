package handlers

import (
	"al-mosso-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

func RemoveFoodFromAppointmentHandler(ctx *fiber.Ctx) error {

	pin := ctx.Params("pin")
	id := ctx.Locals("userId").(uint64)

	var req input
	ctx.BodyParser(&req)

	food_id := req.FoodId

	svc := services.NewRemoveFoodFromAppointmentService()
	err := svc.Execute(pin, id, food_id)

	if err != nil {
		return DispatchError(ctx, *err)
	}

	return NoContent(ctx)
}

type input struct {
	FoodId uint64 `json:"food_id"`
}
