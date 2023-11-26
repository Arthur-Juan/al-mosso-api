package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/interfaces"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
	"errors"
)

type AddFoodToAppointmentService struct {
	db interfaces.IDatabase
}

func NewAddFoodToAppointmentService() *AddFoodToAppointmentService {
	return &AddFoodToAppointmentService{
		db: config.GetDb(),
	}
}

func (s *AddFoodToAppointmentService) Execute(input *types.AddFoodToAppointmentInput, clientId uint64) *error.TError {

	var appointment schemas.Appointment

	if err := s.db.Where("pin = ?", input.PIN).First(&appointment).Error; err != nil {

		if err.Error() == "record not found" {

			return error.NewError(404, errors.New("reserva não encontrada com esse pin"))
		}
		return error.NewError(500, err)
	}

	var client schemas.Client
	if err := s.db.Where("id = ?", clientId).First(&client).Error; err != nil {

		if err.Error() == "record not found" {
			return error.NewError(404, errors.New("usuário não encontrado"))
		}
		return error.NewError(500, err)
	}

	if uint64(client.ID) != appointment.ClientID {
		return error.NewError(403, errors.New("não autorizado a editar essa reserva"))
	}

	var food schemas.Food
	if err := s.db.Where("id = ?", input.FoodId).First(&food).Error; err != nil {

		if err.Error() == "record not found" {
			return error.NewError(404, errors.New("comida não encontrada"))
		}
		return error.NewError(500, err)
	}

	appointment.Foods = append(appointment.Foods, &food)
	appointment.CalculatePrice()
	s.db.Save(&appointment)
	return nil
}
