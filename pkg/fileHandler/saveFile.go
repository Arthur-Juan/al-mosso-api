package fileHandler

import (
	"al-mosso-api/config"
	"al-mosso-api/internal/handler/types"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func SaveFile(file *types.TFile) (string, error) {
	//_, err := os.Stat(config.GetFilePath())
	//if err != nil {
	//	return "", err
	//}
	//if os.IsNotExist(err) {
	//	err = os.MkdirAll(config.GetFilePath(), os.ModePerm)
	//	if err != nil {
	//		return "", err
	//	}
	//
	//	file, err := os.Create(config.GetFilePath())
	//	if err != nil {
	//		return "", err
	//	}
	//
	//	err = file.Close()
	//	if err != nil {
	//		return "", err
	//	}
	//
	//}
	fileName := generateUniqueFileName(file)
	filePath := fmt.Sprintf("%s/%s", config.GetFilePath(), fileName)
	dst, err := os.Create(filePath)

	defer dst.Close()

	if err != nil {
		return "", err
	}
	_, err = io.Copy(dst, file.FileData)
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
