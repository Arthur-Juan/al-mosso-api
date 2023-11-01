package handlers

import (
	"al-mosso-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

func GetAppointmentHandler(ctx *fiber.Ctx) error {

	pin := ctx.Params("pin")
	id := ctx.Locals("userId").(uint64)

	svc := services.NewGetAppointmentService()
	result, err := svc.Execute(pin, id)
	if err != nil {
		return InternalServerError(ctx, err)
	}

	return Ok(ctx, result)
}
