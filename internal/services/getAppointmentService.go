package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
	"errors"

	"gorm.io/gorm"
)

type GetAppointmentService struct {
	db *gorm.DB
}

func NewGetAppointmentService() *GetAppointmentService {
	return &GetAppointmentService{
		db: config.GetDb(),
	}
}

func (s *GetAppointmentService) Execute(pin string, userid uint64) (*types.AppointmentDetailOutput, *error.TError) {

	//verificar se user é dono do PIN (pegar pin e dps user)

	var appointment *schemas.Appointment

	if err := s.db.Preload("Foods").Where("pin = ? and verified = true", pin).First(&appointment).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, error.NewError(404, errors.New("no appointent with this pin"))
		}
		return nil, error.NewError(500, err)
	}

	if appointment.ClientID != userid {
		logger.Infof("userId: %s | appointment.ClientId: %s", userid, appointment.ClientID)
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
	}

	logger.Debug(result)
	return result, nil
}
