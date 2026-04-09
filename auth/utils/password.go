package utils

import (
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, logger *zap.Logger) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("password hashing failed", zap.Error(err))
		return "", fmt.Errorf("password hashing failed: %w", err)
	}
	return string(hash), nil
}

func CheckPassword(hashedPassword, password string, logger *zap.Logger) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		logger.Debug("password comparison failed", zap.Error(err))
		return fmt.Errorf("invalid password")
	}
	return nil
}
