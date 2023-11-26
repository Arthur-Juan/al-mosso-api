package schemas

import (
	"al-mosso-api/internal/entity"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	entity.Appointment
	Foods []*Food `gorm:"many2many:appointment_foods"`
}

func (a *Appointment) CalculatePrice() {
	var val float64

	if a.Foods != nil {
		for _, value := range a.Foods {
			val += value.Price
		}
	}

	val += 25.0 //taxa da reserva
	a.Price = val
}

type Client struct {
	gorm.Model
	entity.Client
}

type Food struct {
	gorm.Model
	entity.Food
}

type Chef struct {
	gorm.Model
	entity.Chef
}

type Contact struct {
	gorm.Model
	entity.Contact
}

type Config struct {
	gorm.Model
	entity.Config
}
