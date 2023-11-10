package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/error"
	"al-mosso-api/pkg/database/schemas"
	"errors"

	"gorm.io/gorm"
)

type RemoveFoodFromAppointmentService struct {
	db *gorm.DB
}

func NewRemoveFoodFromAppointmentService() *RemoveFoodFromAppointmentService {
	return &RemoveFoodFromAppointmentService{
		db: config.GetDb(),
	}
}

func (s *RemoveFoodFromAppointmentService) Execute(pin string, userId uint64, foodId uint64) *error.TError {

	var client *schemas.Client
	if err := s.db.Where("id = ?", userId).First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return error.NewError(404, errors.New("user not found"))
		}
		return error.NewError(500, err)
	}

	var appointment *schemas.Appointment

	if err := s.db.Where("pin = ? and client_id = ? and verified = true", pin, client.ID).First(&appointment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return error.NewError(404, errors.New("appointment not found"))
		}
		return error.NewError(500, err)
	}

	var food *schemas.Food
	if err := s.db.Where("id = ?", foodId).First(&food).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return error.NewError(404, errors.New("food not found"))
		}
		return error.NewError(500, err)
	}

	s.db.Select("Appointments").Where("pin = ? and client_id = ? and verified = true", pin, client.ID).Delete(food).Limit(1)

	return nil
}
