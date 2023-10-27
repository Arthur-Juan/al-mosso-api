package handlers

import (
	"al-mosso-api/internal/services"
	"al-mosso-api/internal/services/types"
	"github.com/gofiber/fiber/v2"
)

func AuthenticateAppointment(ctx *fiber.Ctx) error {

	input := &types.LoginAppointmentInput{}

	err := ctx.BodyParser(&input)
	if err != nil {
		return err
	}

	svc := services.NewAuthenticateAppointmentService()
	result, err := svc.Execute(input)
	if err != nil {
		return InternalServerError(ctx, err)
	}

	response := fiber.Map{
		"token": result,
	}
	return Ok(ctx, response)

}
