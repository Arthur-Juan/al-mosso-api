package cryptography

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func GenerateRandomPassowrd() string {
	length := rand.Intn(15) + 1
	rand.Seed(time.Now().UnixNano())

	// Caracteres permitidos na senha
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#"

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}

func Encrypt(msg string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(msg), 14)
	return string(bytes), err
}
