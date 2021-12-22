package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
