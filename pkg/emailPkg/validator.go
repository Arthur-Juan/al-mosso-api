package emailPkg

import (
	"net/mail"
)

func Validate(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
