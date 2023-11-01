package types

import (
	"al-mosso-api/internal/entity"
	"time"
)

type MakeAppointmentInput struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Date     time.Time `json:"date"`
	Start    string    `json:"start"`
	End      string    `json:"end"`
	Quantity int       `json:"quantity"`
	Message  string    `json:"message"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AppointmentOutput struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AppointmentConfirmation struct {
}

type LoginAppointmentInput struct {
	PIN      string `json:"PIN"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AppointmentDetailOutput struct {
	Date      time.Time
	Start     time.Time
	End       time.Time
	PeopleQtd int
	Message   string
	Foods     []*entity.Food
	PIN       string
}
