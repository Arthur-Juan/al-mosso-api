package entity

import "testing"

func TestNewFoodWithBlankName(t *testing.T) {
	_, err := NewFood("", 123.0, "aaaa")
	if err != nil {
		return
	}
	t.Error("FAIL")
}

func TestNewFoodWithBlankDescription(t *testing.T) {
	_, err := NewFood("apple", 123.0, "")
	if err != nil {
		return
	}
	t.Error("FAIL")
}

func TestNewFoodWithInvalidPrice(t *testing.T) {
	_, err := NewFood("apple", -1, "aa")
	if err != nil {
		return
	}
	t.Error("FAIL")
}

func TestNewFood(t *testing.T) {
	_, err := NewFood("apple", 1, "aa")
	if err == nil {
		return
	}
	t.Error(err)
}
