package handlers

import (
	"al-mosso-api/internal/services"
	"al-mosso-api/internal/services/types"
	"github.com/gofiber/fiber/v2"
	"path/filepath"
)

func InsertChefHandler(ctx *fiber.Ctx) error {
	formFile, err := ctx.FormFile("picture")
	if err != nil {

	}

	file := &types.TFile{}
	input := &types.InsertChefInput{}

	err = ctx.BodyParser(&input)
	if err != nil {

	}

	if formFile != nil {
		file.FileType = formFile.Header.Get("Content-Type")
		file.Extension = filepath.Ext(formFile.Filename)
		fileReader, _ := formFile.Open()
		file.FileData = fileReader
	}

	input.Photo = file
	svc := services.NewInsertChefService()
	result, err := svc.Execute(input)
	if err != nil {

	}
	return Created(ctx, result)
}
