package cryptography

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	random "math/rand"
	"time"
)

func GenerateRandomHash() (string, error) {
	// Generate a random byte slice
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Create a SHA-256 hash of the random data
	hasher := sha256.New()
	_, err = hasher.Write(randomBytes)
	if err != nil {
		return "", err
	}

	// Get the hash as a hexadecimal string
	hashBytes := hasher.Sum(nil)
	hash := hex.EncodeToString(hashBytes)

	return hash, nil
}

func GenerateDecoratedCode() string {
	length := 13
	// Define the decoration pattern
	pattern := "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

	// Initialize a random seed
	random.Seed(time.Now().UnixNano())

	// Generate a random code
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = pattern[random.Intn(len(pattern))]
	}

	return string(code)
}
