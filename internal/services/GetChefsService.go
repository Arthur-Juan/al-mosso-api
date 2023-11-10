package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"

	"gorm.io/gorm"
)

type GetChefsSerivce struct {
	db *gorm.DB
}

func NewGetChefsService() *GetChefsSerivce {
	return &GetChefsSerivce{
		db: config.GetDb(),
	}
}

func (s *GetChefsSerivce) Execute() ([]types.ChefOutput, *error.TError) {
	var chefs []schemas.Chef
	if err := s.db.Find(&chefs).Error; err != nil {
		return nil, error.NewError(500, err)
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
