package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func IsValid(token string, tokenType string) (string, error) {
	var secretKey []byte = []byte(os.Getenv("SECRET_KEY"))
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims["type"] != tokenType {
		return "", fmt.Errorf("Not %v token", tokenType)
	}

	return claims["username"].(string), nil
}

func GenerateToken(username string, tokenType string, d time.Duration) (string, error) {
	var secretKey []byte = []byte(os.Getenv("SECRET_KEY"))
	claims := jwt.MapClaims{
		"username": username,
		"type":     tokenType,
		"exp":      time.Now().Add(d).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
