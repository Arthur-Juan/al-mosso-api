package types

import (
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
