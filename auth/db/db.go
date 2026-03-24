package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// ——————————————————————————————————————————————————————————————————————————————

type InsertRequest struct {
	Username string
	Hash     string
}

// ——————————————————————————————————————————————————————————————————————————————

func ConnectDB() error {
	var connStr string = os.Getenv("DATABASE_URL")

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии базы данных: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Не удалось пингануть бд: %w", err)
	}

	return nil
}

func CloseDB() {
	db.Close()
}

// ——————————————————————————————————————————————————————————————————————————————

func InsertIntoAuth(req *InsertRequest) error {
	_, err := db.Exec("INSERT INTO auth (username, hash) VALUES ($1, $2);", req.Username, req.Hash)
	if err != nil {
		return fmt.Errorf("Не удалось записать в бд: %w", err)
	}

	return nil
}

func DeleteFromAuth(username string) error {
	_, err := db.Exec("Delete from auth where username=$1;", username)
	if err != nil {
		return fmt.Errorf("Ошибка при попытке удаления: %w", err)
	}

	return nil
}

func SelectHash(username string) (string, error) {
	var hash string
	err := db.QueryRow("Select hash from auth where username=$1;", username).Scan(&hash)
	if err != nil {
		return "", fmt.Errorf("Не удалось получить хеш: %w", err)
	}

	return hash, nil
}
