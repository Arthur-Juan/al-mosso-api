package schemas

import (
	"al-mosso-api/internal/entity"
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	entity.Appointment
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
