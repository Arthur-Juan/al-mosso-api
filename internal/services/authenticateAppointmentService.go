package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/cryptography"
	"al-mosso-api/pkg/database/schemas"
	token2 "al-mosso-api/pkg/token"
	"gorm.io/gorm"
)

type AuthenticateAppointment struct {
	db *gorm.DB
}

func NewAuthenticateAppointmentService() *AuthenticateAppointment {
	return &AuthenticateAppointment{
		db: config.GetDb(),
	}
}

func (s *AuthenticateAppointment) Execute(input *types.LoginAppointmentInput) (string, error) {

	var client *schemas.Client

	if err := s.db.Where("email = ?", input.Email).First(&client).Error; err != nil {
		return "", err
	}

	var appointment *schemas.Appointment
	if err := s.db.Where("client_id = ? and pin = ?", client.ID, input.PIN).First(&appointment).Error; err != nil {
		return "", err
	}

	err := cryptography.CheckPassword(input.Password, appointment.Password)
	if err != nil {
		return "", err
	}

	//generate auth token
	token, err := token2.GenerateToken(client)

	return token, nil
}
