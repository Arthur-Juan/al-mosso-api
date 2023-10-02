package entity

import (
	"errors"
	"time"
)

type Appointment struct {
	BaseEntity
	Client    *Client
	ClientID  uint
	Date      time.Time
	Period    time.Time
	PeopleQtd int
	Message   string
	Foods     []*Food `gorm:"many2many:appointment_foods"`
	Verified  bool
}

func NewAppointment(client *Client, date time.Time, period time.Time, peopleQtd int, message string) (*Appointment, error) {

	if client == nil || date.IsZero() || period.IsZero() || peopleQtd < 1 || message == "" {
		return nil, errors.New("client, date, period, people quantity and message is require")
	}

	return &Appointment{
		Client:    client,
		ClientID:  client.ID,
		Date:      date,
		Period:    period,
		PeopleQtd: peopleQtd,
		Message:   message,
	}, nil
}
