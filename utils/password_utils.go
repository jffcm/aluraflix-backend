package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash the password
func HashPassword(password string) (string, error) {
	const cost = 14

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Validate the password
func ValidatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
