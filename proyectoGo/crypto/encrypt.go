package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptText(text string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), cost)
	return string(bytes), err
}
