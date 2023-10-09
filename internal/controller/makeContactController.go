package controller

import (
	"al-mosso-api/internal/handler"
	"al-mosso-api/internal/handler/types"
	"github.com/gofiber/fiber/v2"
)

func MakeContactController(ctx *fiber.Ctx) error {

	input := &types.MakeContactInput{}

	err := ctx.BodyParser(input)
	if err != nil {
		return err
	}
	result, err := handler.MakeContactHandler(input)
	if err != nil {
		return InternalServerError(ctx, err)
	}

	return Created(ctx, result)
}
