package handler

import (
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/handler/types"
	"al-mosso-api/pkg/database/schemas"
	"al-mosso-api/pkg/fileHandler"
	"fmt"
	"gorm.io/gorm"
)

func InsertFoodHandler(input *types.InsertFoodInput) (uint, error) {

	err := input.Validate()
	if err != nil {
		return 0, err
	}

	fmt.Println(input.File.FileData)

	foodEntity, err := entity.NewFood(input.Name, input.Price, input.Description)
	if err != nil {
		return 0, err
	}
	if input.File != nil {
		file, err := fileHandler.SaveFile(input.File)
		if err != nil {
			return 0, err
		}
		foodEntity.ProfilePic = file
	}

	schema := schemas.Food{
		Model: gorm.Model{},
		Food:  *foodEntity,
	}
	result := db.Create(&schema)
	if result.Error != nil {
		return 0, result.Error
	}

	return schema.ID, nil
}
