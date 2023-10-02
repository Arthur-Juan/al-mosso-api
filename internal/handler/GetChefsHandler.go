package handler

import (
	"al-mosso-api/internal/handler/types"
	"al-mosso-api/pkg/database/schemas"
)

func GetChefsHandler() ([]types.ChefOutput, error) {
	var chefs []schemas.Chef
	if err := db.Find(&chefs).Error; err != nil {
		return nil, err
	}
	logger.Infof("Data: %v: ", chefs)
	var result []types.ChefOutput

	for _, chef := range chefs {
		result = append(result, types.ChefOutput{
			ID:          chef.ID,
			Role:        chef.Role,
			Description: chef.Description,
		})
	}

	return result, nil
}
