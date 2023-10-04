package types

import "errors"

type InsertFoodInput struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	File        *File
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
