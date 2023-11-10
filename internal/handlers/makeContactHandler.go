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
	result, terr := svc.Execute(input)
	if terr != nil {
		return DispatchError(ctx, *terr)
	}

	return Created(ctx, result)
}
