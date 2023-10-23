package handlers

import (
	"al-mosso-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

func ConfirmAppointmentHandler(ctx *fiber.Ctx) error {
	hash := ctx.Params("hash")
	svc := services.NewConfirmAppointmentService()
	appointment, err := svc.Execute(hash)
	if err != nil {
		return InternalServerError(ctx, err)
	}

	return Ok(ctx, appointment)
}
