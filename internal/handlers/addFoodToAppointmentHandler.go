package handlers

import (
	"al-mosso-api/internal/services"
	"al-mosso-api/internal/services/types"

	"github.com/gofiber/fiber/v2"
)

func AddFoodToAppointmentHandler(ctx *fiber.Ctx) error {

	var input *types.AddFoodToAppointmentInput
	ctx.BodyParser(&input)

	id := ctx.Locals("userId").(uint64)

	svc := services.NewAddFoodToAppointmentService()
	err := svc.Execute(input, id)

	if err != nil {
		return InternalServerError(ctx, err)
	}

	return NoContent(ctx)
}
