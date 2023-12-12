package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash use bcrypt encrypt the password
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck Compare plaintext passwords to database hashes
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
