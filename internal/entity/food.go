package entity

import "errors"

type Food struct {
	BaseEntity
	Name        string
	Price       float64
	Description string
	ProfilePic  string
}

func NewFood(name string, price float64, description string) (*Food, error) {

	if price < 1 {
		return nil, errors.New("price must be greater than 1")
	}
	if name == "" || price < 1 || description == "" {
		return nil, errors.New("name, price and description fields are required")
	}

	return &Food{
		Name:        name,
		Price:       price,
		Description: description,
	}, nil
}
