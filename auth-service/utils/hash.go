package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword - .
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 5)

	return string(bytes)
}

// CheckPasswordHash - .
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
