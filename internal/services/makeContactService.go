package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"

	"gorm.io/gorm"
)

type MakeContactService struct {
	db *gorm.DB
}

func NewMakeContactService() *MakeContactService {
	return &MakeContactService{
		db: config.GetDb(),
	}
}

func (s *MakeContactService) Execute(input *types.MakeContactInput) (string, *error.TError) {

	contactEntity, err := entity.NewContact(input.Name, input.Email, input.Subject, input.Message)
	if err != nil {
		return "", error.NewError(500, err)
	}

	schema := &schemas.Contact{
		Model:   gorm.Model{},
		Contact: *contactEntity,
	}

	result := s.db.Create(&schema)
	if result.Error != nil {
		return "", error.NewError(500, err)
	}

	return "Contact send!", nil
}
