package handlers

import (
	"al-mosso-api/internal/services"
	"al-mosso-api/internal/services/types"

	"github.com/gofiber/fiber/v2"
)

func AddFoodToAppointmentByTokenHandler(ctx *fiber.Ctx) error {

	var input *types.AddFoodToAppointmentInput
	ctx.BodyParser(&input)

	id := ctx.Locals("userId").(uint64)
	pin := ctx.Locals("pin").(string)
	input.PIN = pin

	svc := services.NewAddFoodToAppointmentService()
	err := svc.Execute(input, id)

	if err != nil {
		return DispatchError(ctx, *err)
	}

	return NoContent(ctx)
}
