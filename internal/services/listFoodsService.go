package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
	"gorm.io/gorm"
)

type ListFoodsService struct {
	db *gorm.DB
}

func NewListFoodsService() *ListFoodsService {
	return &ListFoodsService{
		db: config.GetDb(),
	}
}

func (s *ListFoodsService) Exceute() ([]types.FoodOutput, error) {

	var foods []schemas.Food

	if err := s.db.Find(&foods).Error; err != nil {
		return nil, err
	}

	var result []types.FoodOutput
	for _, food := range foods {
		result = append(result, types.FoodOutput{
			ID:          food.ID,
			Name:        food.Name,
			Price:       food.Price,
			Description: food.Description,
			File:        food.ProfilePic,
		})
	}

	return result, nil
}
