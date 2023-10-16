package controller

import (
	"al-mosso-api/internal/handler"
	"al-mosso-api/internal/handler/types"
	"github.com/gofiber/fiber/v2"
)

func MakeAppointmentController(ctx *fiber.Ctx) error {

	input := &types.MakeAppointmentInput{}

	err := ctx.BodyParser(&input)
	if err != nil {
		return err
	}

	result, err := handler.MakeAppointmentHandler(input)

	return Created(ctx, result)

}
