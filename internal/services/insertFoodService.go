package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
	"al-mosso-api/pkg/fileHandler"
	"fmt"
	"gorm.io/gorm"
)

type InsertFoodService struct {
	db *gorm.DB
}

func NewInsertFoodService() *InsertFoodService {
	return &InsertFoodService{
		db: config.GetDb(),
	}
}

func (s *InsertFoodService) Execute(input *types.InsertFoodInput) (uint, error) {

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
	result := s.db.Create(&schema)
	if result.Error != nil {
		return 0, result.Error
	}

	return schema.ID, nil
}