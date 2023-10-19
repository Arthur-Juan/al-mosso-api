package handlers

import (
	"al-mosso-api/internal/services"
	"al-mosso-api/internal/services/types"
	"github.com/gofiber/fiber/v2"
)

func MakeAppointmentHandler(ctx *fiber.Ctx) error {

	input := &types.MakeAppointmentInput{}

	err := ctx.BodyParser(&input)
	if err != nil {
		return err
	}

	svc := services.NewMakeAppointmentService()
	result, err := svc.Execute(input)

	return Created(ctx, result)

}
