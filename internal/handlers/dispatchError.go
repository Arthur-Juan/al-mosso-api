package handlers

import (
	terror "al-mosso-api/internal/error"

	"github.com/gofiber/fiber/v2"
)

func DispatchError(ctx *fiber.Ctx, err terror.TError) error {

	switch err.Code {
	case 404:
		return NotFound(ctx, err.Err.Error())

	case 403:
		return Unauthorized(ctx, err.Err)
	case 500:
		return InternalServerError(ctx, err.Err)
	}
	return nil
}
