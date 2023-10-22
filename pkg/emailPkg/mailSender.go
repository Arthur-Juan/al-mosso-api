package emailPkg

import (
	logger2 "al-mosso-api/pkg/logger"
	"bytes"
	"encoding/json"
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
	logger := logger2.NewLogger("mail")
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
		logger.Error(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return err
	}

	fmt.Println(string(body))
	response := Response{}

	if err := json.Unmarshal([]byte(body), &response); err != nil {
		logger.Error(err)
		return err
	}

	if response.Errors != nil {
		logger.Error(response.Errors)
		err, _ := json.Marshal(response.Errors)
		return errors.New(string(err))
	}
	return nil
}

type Response struct {
	Success    bool     `json:"success"`
	MessageIDs []string `json:"message_ids"`
	Errors     []string `json:"errors"`
}
