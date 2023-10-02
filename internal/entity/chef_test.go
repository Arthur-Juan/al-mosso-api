package entity

import "testing"

func TestNewChefWithBlankName(t *testing.T) {
	_, err := NewChef("", "chef", "boos chef", "")
	if err != nil {
		return
	}
	t.Error("FAIL")
}

func TestNewChefWithBlankRole(t *testing.T) {
	_, err := NewChef("Walter", "", "boos chef", "")
	if err != nil {
		return
	}
	t.Error("FAIL")
}
func TestNewChefWithBlankDescription(t *testing.T) {
	_, err := NewChef("Walter", "chef", "", "")
	if err != nil {
		return
	}
	t.Error("FAIL")
}
