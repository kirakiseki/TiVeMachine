package utils

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"user/setup"
)

func Encrypt(password string) (string, error) {
	key, err := scrypt.Key([]byte(password), []byte(setup.Config.PasswordSalt), 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(key), nil
}
