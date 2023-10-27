package token

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/entity"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(client entity.Client) (string, error) {
	claims := jwt.MapClaims{
		"sub":   client.ID,
		"name":  client.Name,
		"email": client.Email,
		"exp":   time.Now().Add(time.Hour * 8).Unix(),
	}

	tokenHandler := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenHandler.SignedString([]byte(config.GetKey()))
	if err != nil {
		return "", err
	}

	return token, nil

}
