package controller

import (
	"al-mosso-api/internal/handler"
	"al-mosso-api/internal/handler/types"
	"github.com/gofiber/fiber/v2"
)

func InsertFoodController(ctx *fiber.Ctx) error {

	formFile, err := ctx.FormFile("picture")
	if err != nil {

	}
	fileReader, _ := formFile.Open()
	file := &types.File{
		FileType: formFile.Header.Get("Content-Type"),
		FileData: fileReader,
	}

	input := &types.InsertFoodInput{}

	err = ctx.BodyParser(input)
	if err != nil {

	}
	input.File = file

	result, err := handler.InsertFoodHandler(input)
	if err != nil {

	}
	return Created(ctx, result)
}
