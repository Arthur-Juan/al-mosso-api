package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/entity"
	"al-mosso-api/pkg/cryptography"
	"al-mosso-api/pkg/database/schemas"
	"al-mosso-api/pkg/emailPkg"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ConfirmAppointmentService struct {
	db *gorm.DB
}

func NewConfirmAppointmentService() *ConfirmAppointmentService {
	return &ConfirmAppointmentService{
		db: config.GetDb(),
	}
}

func (s *ConfirmAppointmentService) Execute(hash string) (*schemas.Appointment, error) {
	var appointment *schemas.Appointment

	if err := s.db.Where("hash = ? and verified = false", hash).First(&appointment).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("essa reserva ja foi confirmada")
		}
		return nil, err
	}

	password := cryptography.GenerateRandomPassowrd()
	hash, err := cryptography.Encrypt(password)
	if err != nil {
		return nil, err
	}

	appointment.Password = password
	appointment.Verified = true

	code := cryptography.GenerateDecorativeCode()

	appointment.Code = code
	logger.Info(appointment)
	result := s.db.Save(&appointment)

	if result.Error != nil {

		return nil, result.Error
	}

	var client entity.Client
	if err := s.db.Where("id = ?", appointment.ClientID).First(&client).Error; err != nil {
		return nil, err
	}
	msg := fmt.Sprintf("Reserva confirmada!! Segue os dados de sua reserva:<br>"+
		"<b>Código:</b> %s"+
		"<b>Login:</b> %s"+
		"<b>Senha</b>: %s", appointment.Code, client.Email, appointment.Password)

	mail, err := emailPkg.NewMailSender(client.Email, "Reserva confirmada!", msg)
	if err != nil {
		return nil, err
	}

	err = mail.Send()
	if err != nil {
		return nil, err
	}
	return appointment, nil

}
