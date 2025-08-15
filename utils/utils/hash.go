package utils

import (
	"crypto/rand"
	"math/big"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRandomID() (int64, error) {
	// Generate random number between 1000000000000000000 and 9223372036854775807 (max int64)
	min := big.NewInt(1000000000000000000) // 19 digits minimum
	max := big.NewInt(9223372036854775807) // max int64 value
	
	// Calculate the range
	rangeValue := big.NewInt(0).Sub(max, min)
	
	// Generate random number in range
	randomValue, err := rand.Int(rand.Reader, rangeValue)
	if err != nil {
		return 0, err
	}
	
	// Add minimum value to get final result
	result := big.NewInt(0).Add(min, randomValue)
	return result.Int64(), nil
}
