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
	ClientID  uint64
	Date      time.Time
	Start     time.Time
	End       time.Time
	PeopleQtd int
	Message   string
	Verified  bool
	Hash      string
	Password  string
	Code      string
	PIN       string
	Price     float64
}

func NewAppointment(client *Client, date string, start string, end string, peopleQtd int, message string) (*Appointment, error) {

	if client == nil || date == "" || start == "" || end == "" || peopleQtd < 1 || message == "" {
		return nil, errors.New("client, date, start, end, people quantity and message are required")
	}

	//TODO -> REVER VALIDAÇÃO
	//if end.Sub(start) > 3 {
	//	return nil, errors.New("max 3h of booking are allowed")
	//}

	parsedDate, err := time.Parse("02-01-2006", date)
	if err != nil {
		return nil, err
	}

	appointment := &Appointment{
		Client:    client,
		ClientID:  client.ID,
		Date:      parsedDate,
		PeopleQtd: peopleQtd,
		Message:   message,
		Verified:  false,
	}

	appointment.SetTime(start, end)

	return appointment, nil
}

func (a *Appointment) SetTime(start string, end string) error {
	s, err := time.Parse("15:04", start)
	if err != nil {
		return errors.New(fmt.Sprintf("erro ao gerenciar tempo: %s", err))
	}
	e, err := time.Parse("15:04", end)

	if err != nil {
		return errors.New(fmt.Sprintf("erro ao gerenciar tempo: %s", err))
	}

	if s.After(e) {
		return errors.New("start time needs to be minor then end")
	}

	if !s.IsZero() {
		a.Start = s
	}
	if !e.IsZero() {
		a.End = e
	}

	return nil

}

func (a *Appointment) CheckOverlap(appointments []Appointment) []Appointment {

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
	a.Hash = hash
}
