package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/interfaces"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
	"errors"
)

type GetAppointmentService struct {
	db interfaces.IDatabase
}

func NewGetAppointmentService() *GetAppointmentService {
	return &GetAppointmentService{
		db: config.GetDb(),
	}
}

func (s *GetAppointmentService) Execute(pin string, userid uint64) (*types.AppointmentDetailOutput, *error.TError) {

	//verificar se user é dono do PIN (pegar pin e dps user)

	var appointment *schemas.Appointment

	if err := s.db.Preload("Foods").Preload("Client").Where("pin = ? and verified = true", pin).First(&appointment).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, error.NewError(404, errors.New("no appointent with this pin"))
		}
		return nil, error.NewError(500, err)
	}

	if appointment.ClientID != userid {
		logger.Infof("userId: %d | appointment.ClientId: %d", userid, appointment.ClientID)
		return nil, error.NewError(403, errors.New("unauthorized"))
	}

	appointment.CalculatePrice()

	logger.Debug(appointment.Message)
	logger.Debug(appointment.PeopleQtd)
	result := &types.AppointmentDetailOutput{
		Date:      appointment.Date,
		End:       appointment.End,
		Start:     appointment.Start,
		PeopleQtd: appointment.PeopleQtd,
		Foods:     appointment.Foods,
		PIN:       appointment.PIN,
		Price:     appointment.Price,
		Message:   appointment.Message,
		Client:    appointment.Client,
	}

	logger.Debug(result)
	return result, nil
}
