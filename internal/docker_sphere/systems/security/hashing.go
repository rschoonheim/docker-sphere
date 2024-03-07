package security

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

// HashPassword - hashes a password using bcrypt
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Error hashing password: ", err)
	}
	return string(hash)
}

// VerifyPassword - verifies a password against a hash
func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		slog.Error("Error verifying password: ", err)
		return false
	}
	return true
}
