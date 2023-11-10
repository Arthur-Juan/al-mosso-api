package handlers

import (
	"al-mosso-api/internal/services"
	"al-mosso-api/internal/services/types"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func InsertChefHandler(ctx *fiber.Ctx) error {
	formFile, err := ctx.FormFile("picture")
	if err != nil {

	}

	file := &types.TFile{}
	input := &types.InsertChefInput{}

	err = ctx.BodyParser(&input)
	if err != nil {
		return InternalServerError(ctx, err)
	}

	if formFile != nil {
		file.FileType = formFile.Header.Get("Content-Type")
		file.Extension = filepath.Ext(formFile.Filename)
		fileReader, _ := formFile.Open()
		file.FileData = fileReader
	}

	input.Photo = file
	svc := services.NewInsertChefService()
	result, terr := svc.Execute(input)
	if terr != nil {
		DispatchError(ctx, *terr)
	}

	return Created(ctx, result)
}
