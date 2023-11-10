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
	result, terr := svc.Execute(input)
	if err != nil {
		return DispatchError(ctx, *terr)
	}

	response := fiber.Map{
		"token": result,
	}
	return Ok(ctx, response)

}
