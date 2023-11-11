package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/interfaces"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
)

type ListFoodsService struct {
	db interfaces.IDatabase
}

func NewListFoodsService() *ListFoodsService {
	return &ListFoodsService{
		db: config.GetDb(),
	}
}

func (s *ListFoodsService) Exceute() ([]types.FoodOutput, *error.TError) {

	var foods []schemas.Food

	if err := s.db.Find(&foods).Error; err != nil {
		return nil, error.NewError(500, err)
	}

	var result []types.FoodOutput
	for _, food := range foods {
		result = append(result, types.FoodOutput{
			ID:          uint64(food.ID),
			Name:        food.Name,
			Price:       food.Price,
			Description: food.Description,
			File:        food.ProfilePic,
		})
	}

	return result, nil
}
