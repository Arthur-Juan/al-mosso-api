package handler

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/handler/types"
	"al-mosso-api/pkg/cryptography"
	"al-mosso-api/pkg/emailPkg"
	"fmt"
)

func MakeAppointmentHandler(input *types.MakeAppointmentInput) (*types.AppointmentOutput, error) {
	/*
		* seguri a abordagem de uma credencial para cada agendamento
		  * envia o input
		  * procura pelo email na base
		  * se não achar => cria client novo
		  * se achar => instancia o client
		  * parte para a criação da reserva
	*/

	/**
	TODO => VERIFICAR DATA, HORÁRIO E SE TEM VAGAS DISPONÍVEIS
	*/
	var client *entity.Client

	err := db.Where("email = ?", input.Email).First(&client).Error
	//if err != nil {
	//	return nil, err
	//}

	if client == nil { // criar nova conta
		newClient, err := entity.NewClient(input.Name, input.Email, input.Phone)
		if err != nil {
			return nil, err
		}
		err = db.Create(&newClient).Error
		if err != nil {
			return nil, err
		}
		client = newClient
	}

	appointment, err := entity.NewAppointment(client, input.Date, input.Period, input.Quantity, input.Message)
	hash, err := cryptography.GenerateRandomHash()
	if err != nil {
		return nil, err
	}

	appointment.Hash = hash
	err = db.Create(&appointment).Error
	if err != nil {
		return nil, err
	}
	accessLink := fmt.Sprintf("%s/appointments/%s", config.GetHostName(), hash)
	mailMsg := fmt.Sprintf("Confirme sua reserva em: %s", accessLink)

	mail, err := emailPkg.NewMailSender(input.Email, "Confirmação de reserva", mailMsg)
	if err != nil {
		return nil, err
	}

	err = mail.Send()
	if err != nil {
		return nil, err
	}

	//criar agendamento
	return &types.AppointmentOutput{
		Success: true,
		Message: fmt.Sprintf("Email enviado para: %s\n", input.Email),
	}, nil
}
