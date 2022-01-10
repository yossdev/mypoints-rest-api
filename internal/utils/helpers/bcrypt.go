package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash func change to var for testing purposes
var Hash = func(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

// ValidateHash func change to var for testing purposes
var ValidateHash = func(secret, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(secret), []byte(hash))
}
