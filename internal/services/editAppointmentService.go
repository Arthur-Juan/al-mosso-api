package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/interfaces"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
	"errors"
	"reflect"
	"time"

	"gorm.io/gorm"
)

type EditAppointmentService struct {
	db interfaces.IDatabase
}

func NewEditAppointmentService() *EditAppointmentService {
	return &EditAppointmentService{
		db: config.GetDb(),
	}
}

func (s *EditAppointmentService) Execute(input *types.UpdateAppointmentInput, pin string, userId uint64) *error.TError {
	var client schemas.Client
	if err := s.db.Where("id = ?", userId).First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return error.NewError(404, errors.New("usuário não encontrado"))
		}
		return error.NewError(500, err)
	}
	logger.Error(input)
	var appointment schemas.Appointment
	if err := s.db.Where("pin = ? and client_id = ? and verified = true", pin, userId).First(&appointment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return error.NewError(404, errors.New("reserva não encontrada"))
		}
		return error.NewError(500, err)
	}

	// Update appointment fields from input if not zero
	inputValue := reflect.ValueOf(input).Elem()
	appointmentValue := reflect.ValueOf(&appointment).Elem()
	logger.Debug(inputValue)
	logger.Debug(appointmentValue)
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

	if input.Start != "" || input.End != "" {
		logger.Debug("DEVERIA ESTAR AQUI PORRRA ")
		err := appointment.SetTime(input.Start, input.End)
		if err != nil {
			terr := error.NewError(500, err)
			logger.Error(err)
			return terr
		}
	}

	date, err := time.Parse("02-01-2006", input.Date)
	if err != nil {
		logger.Error(err)
		return error.NewError(500, err)
	}

	appointment.Date = date
	logger.Debug(appointment)

	if err := s.db.Save(&appointment).Error; err != nil {
		return error.NewError(500, err)
	}

	s.db.Model(&appointment).Update("date", date)

	return nil
}
