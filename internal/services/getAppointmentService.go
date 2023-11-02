package services

import (
	"al-mosso-api/config"
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

func (s *GetAppointmentService) Execute(pin string, userid uint64) (*types.AppointmentDetailOutput, error) {

	//verificar se user é dono do PIN (pegar pin e dps user)

	var appointment *schemas.Appointment

	if err := s.db.Preload("Foods").Where("pin = ? and verified = true", pin).First(&appointment).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("no appointent with this pin")
		}
		return nil, err
	}

	logger.Debug(appointment.End)

	if appointment.ClientID != userid {
		logger.Infof("userId: %s | appointment.ClientId: %s", userid, appointment.ClientID)
		return nil, errors.New("unauthorized")
	}

	return &types.AppointmentDetailOutput{
		Date:      appointment.Date,
		End:       appointment.End,
		Start:     appointment.Start,
		PeopleQtd: appointment.PeopleQtd,
		Foods:     appointment.Foods,
		PIN:       appointment.PIN,
	}, nil
}
