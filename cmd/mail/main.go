package main

import (
	"al-mosso-api/pkg/emailPkg"
	"log"
)

func main() {
	mail, err := emailPkg.NewMailSender("arthurjuan214@gmail.com", "teste", "teste")
	if err != nil {
		panic(err)
	}

	err = mail.Send()
	log.Println(err)
}
