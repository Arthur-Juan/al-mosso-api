package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/interfaces"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/database/schemas"
)

type GetChefsSerivce struct {
	db interfaces.IDatabase
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
			Name:        chef.Name,
			Role:        chef.Role,
			Description: chef.Description,
			Photo:       chef.ProfilePic,
		})
	}

	return result, nil
}
