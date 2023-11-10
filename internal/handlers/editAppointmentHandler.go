package handlers

import (
	"al-mosso-api/internal/services"
	"al-mosso-api/internal/services/types"

	"github.com/gofiber/fiber/v2"
)

func EditAppointmentHandler(ctx *fiber.Ctx) error {
	input := &types.UpdateAppointmentInput{}
	ctx.BodyParser(&input)

	pin := ctx.Params("pin")

	id := ctx.Locals("userId").(uint64)

	svc := services.NewEditAppointmentService()

	err := svc.Execute(input, pin, id)

	if err != nil {
		return DispatchError(ctx, *err)
	}

	return NoContent(ctx)
}
