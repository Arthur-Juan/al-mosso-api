package cryptography

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
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
