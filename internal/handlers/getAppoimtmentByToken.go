package handlers

import (
	"al-mosso-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetAppointmentByToken(ctx *fiber.Ctx) error {
	id := ctx.Locals("userId").(uint64)
	pin := ctx.Locals("pin").(string)

	svc := services.NewGetAppointmentService()
	result, err := svc.Execute(pin, id)
	if err != nil {
		return DispatchError(ctx, *err)
	}

	return Ok(ctx, result)

}
