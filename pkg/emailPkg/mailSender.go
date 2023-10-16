package emailPkg

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type MailSender struct {
	To      string
	Subject string
	Message string
}

func NewMailSender(to string, subject string, message string) (*MailSender, error) {
	err := Validate(to)
	if err != nil {
		return nil, err
	}
	if message == "" {
		return nil, errors.New("message is required")
	}
	return &MailSender{
		To:      to,
		Subject: subject,
		Message: message,
	}, nil
}

func (m *MailSender) Send() error {
	msg := []byte(fmt.Sprintf(`{"from":{"email":"arthurjuan214@gmail.com"},
	"to": [{"email":"%s"}],
	"subject":"%s",
	"text":"%s"
}`, m.To, m.Subject, m.Message))

	host, token := getConfig()
	request, err := http.NewRequest(http.MethodPost, host, bytes.NewBuffer(msg))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{Timeout: 5 * time.Second}
	fmt.Println(request.Body)
	res, err := client.Do(request)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
