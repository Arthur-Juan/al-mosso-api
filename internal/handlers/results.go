package handlers

import (
	logger2 "al-mosso-api/pkg/logger"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	logger = logger2.NewLogger("results")
)

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(data)
}

func InternalServerError(ctx *fiber.Ctx, err error) error {
	logger.Errorf("%s", err)
	return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
}

func NotFound(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(http.StatusNotFound).JSON(msg)
}

func Created(ctx *fiber.Ctx, msg interface{}) error {
	return ctx.Status(http.StatusCreated).JSON(msg)
}

func NoContent(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusNoContent)
}

func Unauthorized(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusUnauthorized).JSON(err)
}
