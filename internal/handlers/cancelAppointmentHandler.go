package handlers

import (
	"al-mosso-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

func CancelAppointmentHandler(ctx *fiber.Ctx) error {

	pin := ctx.Params("pin")
	id := ctx.Locals("userId").(uint64)

	svc := services.NewCancelAppointmentService()

	err := svc.Execute(pin, id)

	if err != nil {
		return DispatchError(ctx, *err)
	}
	return NoContent(ctx)
}
