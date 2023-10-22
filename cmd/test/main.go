package main

import (
	"al-mosso-api/internal/entity"
	"fmt"
	"time"
)

func main() {

	client, _ := entity.NewClient("art", "art@email.com", "22992059200")
	appointment, _ := entity.NewAppointment(client, time.Now(), time.Now().Add(time.Hour*12), time.Now().Add(time.Hour*16), 2, "quero cafe")

	overlap1, err := entity.NewAppointment(client, time.Now(), time.Now().Add(time.Hour*17), time.Now().Add(time.Hour*19), 2, "quero churros")
	if err != nil {
		fmt.Println(err)
	}
	overlap2, err := entity.NewAppointment(client, time.Now(), time.Now().Add(time.Hour*15), time.Now().Add(time.Hour*18), 2, "quero pica (overlap)")

	overlap3, err := entity.NewAppointment(client, time.Now(), time.Now().Add(time.Hour*20), time.Now().Add(time.Hour*22), 2, "quero pica (sem overlap)")
	if err != nil {
		fmt.Println(err)
	}
	array := []entity.Appointment{*overlap1, *overlap2, *overlap3}
	overlaps := appointment.CheckOverlap(array)
	var msg []string
	for _, value := range overlaps {
		msg = append(msg, value.Message)
	}
	fmt.Println("Overlaps: ", msg)
}
