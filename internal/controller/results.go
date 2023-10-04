package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(data)
}

func InternalServerError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
}

func NotFound(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(http.StatusNotFound).JSON(msg)
}

func Created(ctx *fiber.Ctx, msg interface{}) error {
	return ctx.Status(http.StatusCreated).JSON(msg)
}
