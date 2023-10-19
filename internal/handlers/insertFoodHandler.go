package handlers

import (
	"al-mosso-api/internal/services"
	"al-mosso-api/internal/services/types"
	"github.com/gofiber/fiber/v2"
	"path/filepath"
)

func InsertFoodHandler(ctx *fiber.Ctx) error {

	formFile, err := ctx.FormFile("picture")
	if err != nil {

	}

	var file = &types.TFile{}

	//if formFile != nil {
	//
	//	fileReader, _ := formFile.Open()
	//	file.FileType = formFile.Header.Get("Content-Type")
	//	file.FileData = fileReader
	//	}
	//
	//}else{
	//
	//}

	input := &types.InsertFoodInput{}

	err = ctx.BodyParser(input)
	if err != nil {

	}

	if formFile != nil {
		fileReader, _ := formFile.Open()
		file.FileData = fileReader
		file.FileType = formFile.Header.Get("Content-Type")
		file.Extension = filepath.Ext(formFile.Filename)
	}

	input.File = file

	svc := services.NewInsertFoodService()
	result, err := svc.Execute(input)
	if err != nil {
		return InternalServerError(ctx, err)
	}
	return Created(ctx, result)
}
