package entity

import (
	"testing"
)

func TestNewClientWithBlankName(t *testing.T) {
	client, err := NewClient("", "email@email.com", "pass123", "22992222222")
	if err != nil {
		return
	}
	t.Error(err)
	print(client)
}

func TestNewClientWithBlankEmail(t *testing.T) {
	client, err := NewClient("name", "", "pass123", "22992222222")
	if err != nil {
		return
	}
	t.Error(err)
	print(client)
}

func TestNewClientWithBlankPass(t *testing.T) {
	client, err := NewClient("name", "email@emal.com", "", "22992222222")
	if err != nil {
		return
	}
	t.Error(err)
	print(client)
}
func TestNewClientWithBlankPhone(t *testing.T) {
	client, err := NewClient("name", "email@emal.com", "pass123", "")
	if err != nil {
		return
	}
	t.Error(err)
	print(client)
}

func TestNewClient(t *testing.T) {
	client, err := NewClient("name", "email@emal.com", "pass123", "22992222222")
	if err != nil {
		t.Error(err)
	}
	print(client)
}
