## Make appointment

User pode fazer reserva em um horário e uma data

input:

```go

type MakeAppointmentInput struct {
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Phone    string        `json:"phone"`
	Date     time.Time     `json:"date"`
	Period   time.Duration `json:"period"`
	Quantity int64         `json:"quantity"`
	Message  string        `json:"message"`
}
```

O sistema deverá verificar se tem lugar disponível para aquela quantidade de pessoas q estão reservando

```go
err := CheckDisponibilty(input.Quantity)
```

onde guardar essa informação de quantidade? (.env/db?)

podemos pegar todas as reservas feitas para X dia X hora, somar o quantity e diminuir de input.Quantity, se o resultado for >= 0, ainda tem vaga e continua com a reserva, se não, cancela a execução e retorna erro

confirmação de email:
deverá ser gerado uma hash única para aquele agendamento, onde quem fez receberá algo como
`http://al-mosso/{unique_hash}`, onde poderá confirmar a reserva,
assinando assim o campo `verified` como true, efetivando a reserva

assim que efetivado, o sistema deverá gerar uma senha para essa reserva, para que o dono possa acessar e gerenciar
`https://al-mosso/appointments/{id}`, e ai poder editar ou cancelar a reserva

então uma reserva se parece com isso:

```go
Appointment{
	Id
	Name (nome do cliente)
	Email (email do cliente)
	Hash (hash unico para verificação do email)
	Data (dia q irá ocorrer)
	Periodo (horário de inicio e fim)
	Quantidade (quantas pessoas nessa reserva)
	Message (mensagem da reserva [Sla veio no front])
	Verified (booleano que diz se foi ou n verificado essa reserva)
	Comidas (será possível atrelar comida a uma reserva, mas isso é outra função)
}

após 2 dias, se n for verificado, a reserva é cancelada


```

## Modelagem
cliente (pronto em entity.Client)

appointment (vai ter o FK do cliente)

na hora de fazer uma reserva, se o usuário não tiver uma conta, será preciso criar uma primeiro

* criar SignUp e SignIn de cliente  =( 
* uma "credencial" para cada reserva


* seguir com abordagem do user ter uma conta 
* todo now => create account [v]

ja cria novo user e basico appointment (nome, msg, dia, "periodo", email) e cria a Hash
* todo => validação de dia, horário, e cadeiras vagas
* next: verificação da reserba no link com hash


