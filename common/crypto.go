package common

import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword convert plain string to hashed
func HashPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hashedByte), err
}

func ValidatePasswordHash(password, hassedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hassedPassword), []byte(password))
	return err == nil
}
