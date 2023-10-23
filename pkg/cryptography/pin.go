package cryptography

import (
	logger2 "al-mosso-api/pkg/logger"
	"fmt"
	"math/rand"
	"time"
)

func GenerateDecorativeCode() string {
	rand.Seed(time.Now().UnixNano())
	min := 10000
	max := 99999
	pin := rand.Intn(max-min+1) + min
	pinStr := fmt.Sprintf("%05d", pin)

	logger := logger2.NewLogger("a")
	logger.Warning(pinStr)
	return (pinStr)
}
