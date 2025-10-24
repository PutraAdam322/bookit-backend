package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPassword string, plainPassword []byte) (bool, error) {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false, err
	}

	return true, nil
}

func HashAndSalt(p string) (string, error) {
	pwd := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
