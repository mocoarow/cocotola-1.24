package gateway

import (
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", liberrors.Errorf("GenerateFromPassword: %w", err)
	}

	return string(hash), nil
}

func ComparePasswords(hashedPassword string, plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return false
	}

	return true
}
