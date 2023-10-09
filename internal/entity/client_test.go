package entity

import (
	"testing"
)

func TestNewClientWithBlankName(t *testing.T) {
	client, err := NewClient("", "emailPkg@emailPkg.com", "22992222222")
	if err != nil {
		return
	}
	t.Error(err)
	print(client)
}

func TestNewClientWithBlankEmail(t *testing.T) {
	client, err := NewClient("name", "", "22992222222")
	if err != nil {
		return
	}
	t.Error(err)
	print(client)
}

func TestNewClientWithBlankPass(t *testing.T) {
	client, err := NewClient("name", "emailPkg@emal.com", "22992222222")
	if err != nil {
		return
	}
	t.Error(err)
	print(client)
}
func TestNewClientWithBlankPhone(t *testing.T) {
	client, err := NewClient("name", "emailPkg@emal.com", "")
	if err != nil {
		return
	}
	t.Error(err)
	print(client)
}

func TestNewClient(t *testing.T) {
	client, err := NewClient("name", "emailPkg@emal.com", "22992222222")
	if err != nil {
		t.Error(err)
	}
	print(client)
}
