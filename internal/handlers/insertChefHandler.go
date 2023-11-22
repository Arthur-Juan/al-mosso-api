package handlers

import (
	"al-mosso-api/internal/services"
	"al-mosso-api/internal/services/types"
	"io"
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
		fileReader, _ := formFile.Open()
		defer fileReader.Close()

		buf := make([]byte, 1024)

		for {
			n, err := fileReader.Read(buf)
			if err == io.EOF {
				break
			}

			file.FileData = append(file.FileData, buf[:n]...)
		}
		fileReader.Read(file.FileData)
		file.FileType = formFile.Header.Get("Content-Type")
		file.Extension = filepath.Ext(formFile.Filename)
	}

	input.Photo = file
	svc := services.NewInsertChefService()
	result, terr := svc.Execute(input)
	if terr != nil {
		DispatchError(ctx, *terr)
	}

	return Created(ctx, result)
}
