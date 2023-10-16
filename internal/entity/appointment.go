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
	Start     time.Time
	End       time.Time
	PeopleQtd int
	Message   string
	Foods     []*Food `gorm:"many2many:appointment_foods"`
	Verified  bool
	Hash      string
	Password  string
}

func NewAppointment(client *Client, date time.Time, start time.Time, end time.Time, peopleQtd int, message string) (*Appointment, error) {

	if client == nil || date.IsZero() || start.IsZero() || end.IsZero() || peopleQtd < 1 || message == "" {
		return nil, errors.New("client, date, period, people quantity and message is require")
	}

	if start.After(end) {
		return nil, errors.New("start time needs to be minor then end")
	}

	if end.Sub(start) > 3 {
		return nil, errors.New("max 3h of booking are allowed")
	}
	return &Appointment{
		Client:    client,
		ClientID:  client.ID,
		Date:      date,
		Start:     start,
		End:       end,
		PeopleQtd: peopleQtd,
		Message:   message,
		Verified:  false,
	}, nil
}
