package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func ValidateHash(secret, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(secret), []byte(hash))
}
