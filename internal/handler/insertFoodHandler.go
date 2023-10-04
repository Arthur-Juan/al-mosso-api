package handler

import (
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/handler/types"
	"al-mosso-api/pkg/database/schemas"
	"gorm.io/gorm"
)

func InsertFoodHandler(input *types.InsertFoodInput) (uint, error) {

	err := input.Validate()
	if err != nil {
		return 0, err
	}

	foodEntity, err := entity.NewFood(input.Name, input.Price, input.Description)
	if err != nil {
		return 0, err
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
