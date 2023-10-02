package entity

import "errors"

type Client struct {
	BaseEntity
	Name         string
	Email        string
	Password     string
	Phone        string
	Appointments []*Appointment
}

func NewClient(name string, email string, phone string) (*Client, error) {

	if name == "" || email == "" || phone == "" {
		return nil, errors.New("all fields are required")
	}

	return &Client{
		Name:  name,
		Email: email,
		Phone: phone,
	}, nil
}

func (client *Client) AddAppointment(appointment *Appointment) error {
	if appointment == nil {
		return errors.New("send an appointment")
	}
	_ = append(client.Appointments, appointment)
	return nil
}
