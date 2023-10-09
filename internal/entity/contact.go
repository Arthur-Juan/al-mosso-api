package entity

import (
	"al-mosso-api/pkg/emailPkg"
	"errors"
)

type Contact struct {
	BaseEntity
	Name    string
	Email   string
	Subject string
	Message string
}

func NewContact(name string, email string, subject string, message string) (*Contact, error) {

	err := emailPkg.Validate(email)

	if err != nil {
		return nil, err
	}

	if name == "" || subject == "" || message == "" {
		return nil, errors.New("name, email, message and subject are required")
	}

	return &Contact{
		Name:    name,
		Email:   email,
		Subject: subject,
		Message: message,
	}, nil
}
