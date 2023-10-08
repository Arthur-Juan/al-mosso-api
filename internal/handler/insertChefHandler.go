package handler

import (
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/handler/types"
	"al-mosso-api/pkg/database/schemas"
	"al-mosso-api/pkg/fileHandler"
	"gorm.io/gorm"
)

func InsertChefHandler(input *types.InsertChefInput) (uint, error) {

	chefEntity, err := entity.NewChef(input.Name, input.Role, input.Description, "")
	if err != nil {
		return 0, err
	}

	if input.Photo != nil {
		file, err := fileHandler.SaveFile(input.Photo)
		if err != nil {
			return 0, err
		}
		chefEntity.ProfilePic = file
	}

	schema := &schemas.Chef{
		Model: gorm.Model{},
		Chef:  *chefEntity,
	}

	result := db.Create(&schema)
	if result.Error != nil {
		return 0, err
	}

	return schema.ID, nil

	return 0, nil
}
