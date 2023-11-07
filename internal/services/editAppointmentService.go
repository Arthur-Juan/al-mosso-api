package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
	"errors"
	"reflect"

	"gorm.io/gorm"
)

type EditAppointmentService struct {
	db *gorm.DB
}

func NewEditAppointmentService() *EditAppointmentService {
	return &EditAppointmentService{
		db: config.GetDb(),
	}
}

func (s *EditAppointmentService) Execute(input *types.UpdateAppointmentInput, pin string, userId uint64) error {
	var client schemas.Client
	if err := s.db.Where("id = ?", userId).First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("usuário não encontrado")
		}
		return err
	}

	var appointment schemas.Appointment
	if err := s.db.Where("pin = ? and client_id = ? and verified = true", pin, userId).First(&appointment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("reserva não encontrada")
		}
		return err
	}

	// Update appointment fields from input if not zero
	inputValue := reflect.ValueOf(input).Elem()
	appointmentValue := reflect.ValueOf(&appointment).Elem()

	for i := 0; i < inputValue.NumField(); i++ {
		inputField := inputValue.Field(i)
		fieldName := inputValue.Type().Field(i).Name

		// Check if the input field is non-zero
		if !inputField.IsZero() {
			appointmentField := appointmentValue.FieldByName(fieldName)
			if appointmentField.IsValid() && appointmentField.Type().AssignableTo(inputField.Type()) {
				appointmentField.Set(inputField)
			}
		}
	}

	logger.Debug(appointment)

	if err := s.db.Save(&appointment).Error; err != nil {
		return err
	}

	return nil
}
