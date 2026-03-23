package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Возникла проблема в момент хеширования: %w", err)
	}
	return string(hash), nil
}

func CheckPassword(hashedPassword, password string) error {
	var err error = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("Ошибка при сверки пароля и хеша: %w", err)
	}

	return nil
}
