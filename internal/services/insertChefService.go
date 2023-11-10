package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
	"al-mosso-api/pkg/fileHandler"

	"gorm.io/gorm"
)

type InsertChefService struct {
	db *gorm.DB
}

func NewInsertChefService() *InsertChefService {
	return &InsertChefService{
		db: config.GetDb(),
	}
}

func (s *InsertChefService) Execute(input *types.InsertChefInput) (uint, *error.TError) {

	chefEntity, err := entity.NewChef(input.Name, input.Role, input.Description, "")
	if err != nil {
		return 0, error.NewError(500, err)
	}

	if input.Photo != nil {
		file, err := fileHandler.SaveFile(input.Photo)
		if err != nil {
			return 0, error.NewError(500, err)
		}
		chefEntity.ProfilePic = file
	}

	schema := &schemas.Chef{
		Model: gorm.Model{},
		Chef:  *chefEntity,
	}

	result := s.db.Create(&schema)
	if result.Error != nil {
		return 0, error.NewError(500, err)
	}

	return schema.ID, nil
}
