package fileHandler

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/services/types"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func SaveFile(file *types.TFile) (string, error) {

	isValid := CheckFile(file)
	if !isValid {
		return "", errors.New("invalid file!")
	}

	fileName := generateUniqueFileName(file)
	filePath := fmt.Sprintf("%s/%s", config.GetFilePath(), fileName)
	dst, err := os.Create(filePath)
	defer dst.Close()

	if err != nil {
		return "", err
	}

	dst.Write(file.FileData)

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", config.GetFilePath(), fileName), nil
}

func generateUniqueFileName(file *types.TFile) string {
	// Get the current timestamp in nanoseconds
	timestamp := time.Now().UnixNano()

	// Generate a random number between 10000 and 99999
	randomNum := rand.Intn(90000) + 10000

	// Combine the timestamp and random number to create a unique filename
	uniqueFileName := fmt.Sprintf("%d%d%s", timestamp, randomNum, file.Extension)

	return uniqueFileName
}
