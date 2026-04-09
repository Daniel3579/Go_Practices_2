package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var db *sql.DB

// ——————————————————————————————————————————————————————————————————————————————

type InsertRequest struct {
	Username string
	Hash     string
}

// ——————————————————————————————————————————————————————————————————————————————

func ConnectDB(logger *zap.Logger) error {
	var connStr string = os.Getenv("DATABASE_URL")

	if connStr == "" {
		logger.Error("DATABASE_URL environment variable is not set")
		return fmt.Errorf("DATABASE_URL not set")
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed to open database", zap.Error(err))
		return fmt.Errorf("Ошибка при открытии базы данных: %w", err)
	}

	err = db.Ping()
	if err != nil {
		logger.Error("Failed to ping database", zap.Error(err))
		return fmt.Errorf("Не удалось пингануть бд: %w", err)
	}

	logger.Info("Successfully connected to database")
	return nil
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// ——————————————————————————————————————————————————————————————————————————————

func InsertIntoAuth(req *InsertRequest, logger *zap.Logger) error {
	_, err := db.Exec("INSERT INTO auth (username, hash) VALUES ($1, $2);", req.Username, req.Hash)
	if err != nil {
		logger.Error("Failed to insert into auth table",
			zap.String("username", req.Username),
			zap.Error(err),
		)
		return fmt.Errorf("Не удалось записать в бд: %w", err)
	}

	logger.Info("User registered successfully", zap.String("username", req.Username))
	return nil
}

func DeleteFromAuth(username string, logger *zap.Logger) error {
	res, err := db.Exec("Delete from auth where username=$1;", username)
	if err != nil {
		logger.Error("Failed to delete from auth table",
			zap.String("username", username),
			zap.Error(err),
		)
		return fmt.Errorf("Ошибка при попытке удаления: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Error("Failed to get rows affected",
			zap.String("username", username),
			zap.Error(err),
		)
		return fmt.Errorf("Failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		logger.Warn("No rows affected for delete operation", zap.String("username", username))
		return fmt.Errorf("User not found")
	}

	logger.Info("User deleted successfully", zap.String("username", username))
	return nil
}

func SelectHash(username string, logger *zap.Logger) (string, error) {
	var hash string
	err := db.QueryRow("Select hash from auth where username=$1;", username).Scan(&hash)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("User not found", zap.String("username", username))
			return "", fmt.Errorf("User not found")
		}
		logger.Error("Failed to query hash",
			zap.String("username", username),
			zap.Error(err),
		)
		return "", fmt.Errorf("Не удалось получить хеш: %w", err)
	}

	return hash, nil
}
