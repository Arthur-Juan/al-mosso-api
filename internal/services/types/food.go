package types

import "errors"

type InsertFoodInput struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	File        *TFile
}

func (i *InsertFoodInput) Validate() error {
	if i.Name == "" || i.Description == "" {
		return errors.New("name and Description are required")
	}

	if i.Price < 1 {
		return errors.New("price must be greater than zero")
	}

	return nil
}

type FoodOutput struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	File        string  `json:"file"`
}
