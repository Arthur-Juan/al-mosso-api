package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/interfaces"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
	"al-mosso-api/pkg/fileHandler"

	"gorm.io/gorm"
)

type InsertFoodService struct {
	db interfaces.IDatabase
}

func NewInsertFoodService() *InsertFoodService {
	return &InsertFoodService{
		db: config.GetDb(),
	}
}

func (s *InsertFoodService) Execute(input *types.InsertFoodInput) (uint, *error.TError) {

	err := input.Validate()
	if err != nil {
		return 0, error.NewError(500, err)
	}

	foodEntity, err := entity.NewFood(input.Name, input.Price, input.Description)
	if err != nil {
		return 0, error.NewError(500, err)
	}
	if input.File != nil {
		file, err := fileHandler.SaveFile(input.File)
		if err != nil {
			return 0, error.NewError(500, err)
		}
		foodEntity.ProfilePic = file
	}

	schema := schemas.Food{
		Model: gorm.Model{},
		Food:  *foodEntity,
	}
	result := s.db.Create(&schema)
	if result.Error != nil {
		return 0, error.NewError(500, result.Error)
	}

	return schema.ID, nil
}
