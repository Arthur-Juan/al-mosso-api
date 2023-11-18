package services

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/entity"
	"al-mosso-api/internal/error"
	"al-mosso-api/internal/interfaces"
	"al-mosso-api/internal/services/types"
	"al-mosso-api/pkg/cryptography"
	"al-mosso-api/pkg/emailPkg"
	"errors"
	"fmt"
)

type MakeAppointmentService struct {
	db interfaces.IDatabase
}

func NewMakeAppointmentService() *MakeAppointmentService {
	return &MakeAppointmentService{
		db: config.GetDb(),
	}
}

func (s *MakeAppointmentService) Execute(input *types.MakeAppointmentInput) (*types.AppointmentOutput, *error.TError) {
	/*
		* seguir a abordagem de uma credencial para cada agendamento
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

	err := s.db.Where("email = ?", input.Email).First(&client).Error

	if err != nil {
		if err.Error() == "record not found" { // criar nova conta
			newClient, err := entity.NewClient(input.Name, input.Email, input.Phone)
			if err != nil {
				return nil, error.NewError(500, err)
			}
			err = s.db.Create(&newClient).Error
			if err != nil {
				return nil, error.NewError(500, err)
			}
			client = newClient
			logger.Info(client)
			logger.Info(newClient)

		}
	}

	//cria entidade de appointment
	appointment, err := entity.NewAppointment(client, input.Date, input.Start, input.End, input.Quantity, input.Message)

	if err != nil {
		return nil, error.NewError(500, err)
	}
	//check quantidade de vagas
	var appointmentsToCheck []entity.Appointment

	//TODO => pegar somente vagas do mesmo dia !! (e talvez com horário proximo, ex: input.Start -3, input.End -3)
	if err := s.db.Find(&appointmentsToCheck).Error; err != nil {
		return nil, error.NewError(500, err)
	}

	overlaps := appointment.CheckOverlap(appointmentsToCheck)

	var overlapsQtd int
	for _, value := range overlaps {
		overlapsQtd += value.PeopleQtd
	}
	if (config.GetVacancies() - overlapsQtd) < input.Quantity {
		return nil, error.NewError(500, errors.New(fmt.Sprintf("não temos vagas suficiente para %d pessoas nesse horário", input.Quantity)))
	}

	//gera hash e salva no banco
	hash, err := cryptography.GenerateRandomHash()
	if err != nil {
		return nil, error.NewError(500, err)
	}

	appointment.SetHash(hash)
	err = s.db.Create(&appointment).Error
	if err != nil {
		return nil, error.NewError(500, err)
	}

	//link de ativação
	accessLink := fmt.Sprintf("%s/api/v1/appointments/confirm/%s", config.GetHostName(), hash)
	mailMsg := fmt.Sprintf("Confirme sua reserva em: %s", accessLink)

	mail, err := emailPkg.NewMailSender(input.Email, "Confirmação de reserva", mailMsg)
	if err != nil {
		return nil, error.NewError(500, err)
	}

	err = mail.Send()
	if err != nil {
		return nil, error.NewError(500, err)
	}

	//criar agendamento
	return &types.AppointmentOutput{
		Success: true,
		Message: fmt.Sprintf("Email enviado para: %s", input.Email),
	}, nil
}
