package handlers

import (
	"al-mosso-api/internal/services"
	"al-mosso-api/internal/services/types"
	"github.com/gofiber/fiber/v2"
)

func MakeContactHandler(ctx *fiber.Ctx) error {

	input := &types.MakeContactInput{}

	err := ctx.BodyParser(input)
	if err != nil {
		return err
	}
	svc := services.NewMakeContactService()
	result, err := svc.Execute(input)
	if err != nil {
		return InternalServerError(ctx, err)
	}

	return Created(ctx, result)
}
