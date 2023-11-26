package token

import (
	"al-mosso-api/config"
	"al-mosso-api/pkg/database/schemas"
	logger2 "al-mosso-api/pkg/logger"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	ID    uint64
	Email string
	Name  string
	Exp   int64
}

func GenerateToken(client *schemas.Client, pin string) (string, error) {

	logger2.NewLogger("auth").Info(client)
	logger2.NewLogger("auth").Info(client.ID)

	claims := &jwt.MapClaims{
		"ID":    uint64(client.ID),
		"Email": client.Email,
		"Name":  client.Email,
		"exp":   time.Now().Add(time.Hour * 8).Unix(),
		"pin":   pin,
	}

	tokenHandler := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenHandler.SignedString([]byte(config.GetKey()))
	if err != nil {
		return "", err
	}

	return token, nil

}

func CheckToken(token string) (Claims, error) {
	var claims Claims
	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetKey()), nil
	})

	if err != nil {
		logger2.NewLogger("auth").Error(err)
		return claims, err
	}
	if !jwtToken.Valid {
		return claims, errors.New("unauthorized")
	}

	return claims, nil
}
