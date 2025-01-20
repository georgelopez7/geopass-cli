package generator

import (
	"crypto/rand"
	"geopass-cli/config"
	"math"
	"math/big"
)

const MinPasswordLength = config.DefaultPasswordLength

const (
	lowercase    = "abcdefghijklmnopqrstuvwxyz"
	uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers      = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
)

func GeneratePassword(length int) (string, float64, error) {

	// --- GENERATE CHARACTER POOL ---
	charPool := lowercase + uppercase + numbers + specialChars

	// --- ENSURE LENGTH IS VALID ---
	if length < MinPasswordLength {
		return "", 0, nil
	}

	// --- GENERATE PASSWORD ---
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charPool))))
		if err != nil {
			return "", 0, err
		}
		password[i] = charPool[idx.Int64()]
	}

	// --- CALCULATE PASSWORD ENTROPY ---
	entropy := calculateEntropy(string(password), len(charPool))

	return string(password), entropy, nil
}

func calculateEntropy(password string, poolSize int) float64 {
	if len(password) == 0 || poolSize <= 1 {
		return 0
	}

	// --- Shannon entropy: H = -Σ (p * log2(p)) ---
	entropyPerChar := math.Log2(float64(poolSize))
	totalEntropy := entropyPerChar * float64(len(password))

	return totalEntropy
}
