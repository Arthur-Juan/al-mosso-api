package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/interfaces"
	"al-mosso-api/internal/services/types"
)

type GetFoodByIdService struct {
	db interfaces.IDatabase
}

func NewGetFoodByIdService() *GetFoodByIdService {
	return &GetFoodByIdService{
		db: config.GetDb(),
	}
}

func (s *GetFoodByIdService) Execute(id uint64) (*types.FoodOutput, *error.TError) {
	var food entity.Food

	if err := s.db.Where("id = ?", id).First(&food).Error; err != nil {
		return nil, error.NewError(500, err)
	}

	result := &types.FoodOutput{
		ID:          food.ID,
		Name:        food.Name,
		Price:       food.Price,
		Description: food.Description,
		File:        food.ProfilePic,
	}
	return result, nil
}
