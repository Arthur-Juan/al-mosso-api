package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/interfaces"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/cryptography"
	"al-mosso-api/pkg/database/schemas"
	token2 "al-mosso-api/pkg/token"
)

type AuthenticateAppointment struct {
	db interfaces.IDatabase
}

func NewAuthenticateAppointmentService() *AuthenticateAppointment {
	return &AuthenticateAppointment{
		db: config.GetDb(),
	}
}

func (s *AuthenticateAppointment) Execute(input *types.LoginAppointmentInput) (string, *error.TError) {

	var client *schemas.Client

	if err := s.db.Where("email = ?", input.Email).First(&client).Error; err != nil {
		return "", error.NewError(404, err)
	}

	var appointment *schemas.Appointment
	if err := s.db.Where("client_id = ? and pin = ?", client.ID, input.PIN).First(&appointment).Error; err != nil {
		return "", error.NewError(404, err)
	}

	err := cryptography.CheckPassword(input.Password, appointment.Password)
	if err != nil {
		return "", error.NewError(500, err)
	}

	//generate auth token
	token, err := token2.GenerateToken(client)

	return token, nil
}
