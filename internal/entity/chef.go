package entity

import "errors"

type Chef struct {
	BaseEntity
	Name        string
	Role        string
	Description string
	ProfilePic  string
}

func NewChef(name string, role string, description string, profilePic string) (*Chef, error) {

	if name == "" || description == "" || role == "" {
		return nil, errors.New("name, role and description are required")
	}

	return &Chef{
		Name:        name,
		Role:        role,
		Description: description,
		ProfilePic:  profilePic,
	}, nil
}
