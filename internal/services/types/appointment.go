package types

import (
	"al-mosso-api/internal/entity"
	"time"
)

type MakeAppointmentInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Date     string `json:"date"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Quantity string `json:"quantity"`
	Message  string `json:"message"`
}

type UpdateAppointmentInput struct {
	PeopleQtd int       `json:"people_quantity"`
	Message   string    `json:"message"`
	Date      time.Time `json:"date"`
	Start     string    `json:"start"`
	End       string    `json:"end"`
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
	Date      time.Time      `json:"date"`
	Start     time.Time      `json:"start"`
	End       time.Time      `json:"end"`
	PeopleQtd int            `json:"people_qtd"`
	Message   string         `json:"message"`
	Foods     []*entity.Food `json:"foods"`
	PIN       string         `json:"PIN"`
	Price     float64        `json:"price"`
}
