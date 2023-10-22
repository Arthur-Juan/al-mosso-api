package entity

import (
	"errors"
	"fmt"
	"sort"
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
	Code      string
}

func NewAppointment(client *Client, date time.Time, start time.Time, end time.Time, peopleQtd int, message string) (*Appointment, error) {

	if client == nil || date.IsZero() || start.IsZero() || end.IsZero() || peopleQtd < 1 || message == "" {
		return nil, errors.New("client, date, start, end, people quantity and message are required")
	}

	if start.After(end) {
		return nil, errors.New("start time needs to be minor then end")
	}

	//TODO -> REVER VALIDAÇÃO
	//if end.Sub(start) > 3 {
	//	return nil, errors.New("max 3h of booking are allowed")
	//}
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

func (a *Appointment) CheckOverlap(appointments []Appointment) []Appointment {

	// Don't create a shallow copy of 'a'
	sort.Slice(appointments, func(i, j int) bool {
		return appointments[i].Start.Before(appointments[j].Start)
	})

	var overlaps []Appointment
	for i := 1; i < len(appointments); i++ {
		if appointments[i-1].End.After(appointments[i].Start) {
			overlaps = append(overlaps, appointments[i])
		}
	}

	return overlaps
}

func (a *Appointment) SetHash(hash string) {
	fmt.Println(hash)
	a.Hash = string(hash)
}
