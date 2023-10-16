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
	Period    string
	PeopleQtd int
	Message   string
	Foods     []*Food `gorm:"many2many:appointment_foods"`
	Verified  bool
	Hash      string
}

func NewAppointment(client *Client, date time.Time, period string, peopleQtd int, message string) (*Appointment, error) {

	if client == nil || date.IsZero() || period == "" || peopleQtd < 1 || message == "" {
		return nil, errors.New("client, date, period, people quantity and message is require")
	}

	return &Appointment{
		Client:    client,
		ClientID:  client.ID,
		Date:      date,
		Period:    period,
		PeopleQtd: peopleQtd,
		Message:   message,
		Verified:  false,
	}, nil
}
