package handlers

import (
	"al-mosso-api/internal/services"
	"al-mosso-api/internal/services/types"
	"io/ioutil"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func InsertChefHandler(ctx *fiber.Ctx) error {
	// Check if the "picture" field is present in the form data
	formFile, err := ctx.FormFile("picture")
	if err != nil {
		// Handle the error, e.g., return BadRequest(ctx, err)
		return BadRequest(ctx, err)
	}

	// Parse other form fields manually
	input := &types.InsertChefInput{
		Name:        ctx.FormValue("name"),
		Role:        ctx.FormValue("role"),
		Description: ctx.FormValue("description"),
	}

	// Create TFile struct to hold file data
	file := &types.TFile{}

	// Check if a file is present in the form data
	if formFile != nil {
		fileReader, err := formFile.Open()
		if err != nil {
			// Handle the error, e.g., return InternalServerError(ctx, err)
			return InternalServerError(ctx, err)
		}
		defer fileReader.Close()

		// Read the entire file data
		buf, err := ioutil.ReadAll(fileReader)
		if err != nil {
			// Handle the error, e.g., return InternalServerError(ctx, err)
			return InternalServerError(ctx, err)
		}

		file.FileData = buf
		file.FileType = formFile.Header.Get("Content-Type")
		file.Extension = filepath.Ext(formFile.Filename)
	}

	// Assign the file data to the input struct
	input.Photo = file

	// Create a service and execute the operation
	svc := services.NewInsertChefService()
	result, terr := svc.Execute(input)
	if terr != nil {
		// Handle the error, e.g., DispatchError(ctx, *terr)
		return DispatchError(ctx, *terr)
	}

	// Return the successful response
	return Created(ctx, result)
}
