package Utills

import (
	"golang.org/x/crypto/bcrypt"
)

func EnCodePassWord(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err

}
