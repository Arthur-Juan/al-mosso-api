package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/interfaces"
	"al-mosso-api/pkg/cryptography"
	"al-mosso-api/pkg/database/schemas"
	"al-mosso-api/pkg/emailPkg"
	"errors"
	"fmt"
)

type ConfirmAppointmentService struct {
	db interfaces.IDatabase
}

func NewConfirmAppointmentService() *ConfirmAppointmentService {
	return &ConfirmAppointmentService{
		db: config.GetDb(),
	}
}

func (s *ConfirmAppointmentService) Execute(hash string) (string, *error.TError) {
	var appointment *schemas.Appointment

	if err := s.db.Where("hash = ? and verified = false", hash).First(&appointment).Error; err != nil {
		if err.Error() == "record not found" {
			return "", error.NewError(404, errors.New("essa reserva ja foi confirmada"))
		}
		return "", error.NewError(500, err)
	}

	password := cryptography.GenerateRandomPassowrd()
	passHash, err := cryptography.Encrypt(password)
	if err != nil {
		return "", error.NewError(500, err)
	}

	appointment.Password = passHash
	appointment.Verified = true

	pin := cryptography.GenerateDecorativeCode()

	appointment.PIN = pin
	logger.Info(appointment)

	result := s.db.Save(&appointment)

	if result.Error != nil {

		return "", error.NewError(500, result.Error)
	}

	var client entity.Client
	if err := s.db.Where("id = ?", appointment.ClientID).First(&client).Error; err != nil {
		return "", error.NewError(500, result.Error)
	}
	logger.Error(client)
	msg := fmt.Sprintf("Reserva confirmada!! Segue os dados de sua reserva:<br>"+
		"<b>Código:</b> %s"+
		"<b>Login:</b> %s"+
		"<b>Senha</b>: %s", appointment.PIN, client.Email, password)

	mail, err := emailPkg.NewMailSender(client.Email, "Reserva confirmada!", msg)
	if err != nil {
		return "", error.NewError(500, err)
	}

	err = mail.Send()
	if err != nil {
		return "", error.NewError(500, err)
	}
	return "Reserva confirmada! Dados enviado para o email", nil

}
