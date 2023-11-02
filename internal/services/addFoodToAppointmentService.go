package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
	"errors"

	"gorm.io/gorm"
)

type AddFoodToAppointmentService struct {
	db *gorm.DB
}

func NewAddFoodToAppointmentService() *AddFoodToAppointmentService {
	return &AddFoodToAppointmentService{
		db: config.GetDb(),
	}
}

func (s *AddFoodToAppointmentService) Execute(input *types.AddFoodToAppointmentInput, clientId uint64) error {

	var appointment schemas.Appointment

	if err := s.db.Where("pin = ?", input.PIN).First(&appointment).Error; err != nil {

		if err.Error() == "record not found" {
			return errors.New("reserva não encontrada com esse pin")
		}
		return err
	}

	var client schemas.Client
	if err := s.db.Where("id = ?", clientId).First(&client).Error; err != nil {

		if err.Error() == "record not found" {
			return errors.New("usuário não encontrado")
		}
		return err
	}

	if uint64(client.ID) != appointment.ClientID {
		return errors.New("não autorizado a editar essa reserva")
	}

	var food schemas.Food
	if err := s.db.Where("id = ?", input.FoodId).First(&food).Error; err != nil {

		if err.Error() == "record not found" {
			return errors.New("comida não encontrada")
		}
		return err
	}

	appointment.Foods = append(appointment.Foods, &food.Food)
	appointment.CalculatePrice()
	s.db.Save(&appointment)
	return nil
}
