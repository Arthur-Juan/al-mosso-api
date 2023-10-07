package handler

import (
	"al-mosso-api/internal/handler/types"
	"al-mosso-api/pkg/database/schemas"
)

func ListFoodsHandler() ([]types.FoodOutput, error) {

	var foods []schemas.Food

	if err := db.Find(&foods).Error; err != nil {
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
