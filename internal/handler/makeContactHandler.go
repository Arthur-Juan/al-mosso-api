package handler

import (
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/handler/types"
	"al-mosso-api/pkg/database/schemas"
	"gorm.io/gorm"
)

func MakeContactHandler(input *types.MakeContactInput) (string, error) {

	contactEntity, err := entity.NewContact(input.Name, input.Email, input.Subject, input.Message)
	if err != nil {
		return "", err
	}

	schema := &schemas.Contact{
		Model:   gorm.Model{},
		Contact: *contactEntity,
	}

	result := db.Create(&schema)
	if result.Error != nil {
		return "", err
	}

	return "Contact send!", nil
}
