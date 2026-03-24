package db

import (
	"database/sql"
	"fmt"
	"os"
	"separation/task/dtos"

	_ "github.com/lib/pq"
)

var db *sql.DB

// ——————————————————————————————————————————————————————————————————————————————

func ConnectDB(env string) error {
	var connStr string = os.Getenv(env)

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

func InsertIntoTask(username string, req *dtos.InsertRequest) (*dtos.SelectResponse, error) {
	var res *dtos.SelectResponse = &dtos.SelectResponse{}

	err := db.QueryRow("INSERT INTO task (username, title, description, due_date) VALUES ($1, $2, $3, $4) RETURNING *;",
		username,
		req.Title,
		req.Description,
		req.Due_date,
	).Scan(&res.Id, &res.Username, &res.Title, &res.Description, &res.Due_date, &res.Done)

	if err != nil {
		return nil, fmt.Errorf("Не удалось записать в бд: %w", err)
	}

	return res, nil
}

func SelectAllTasks(username string) (*[]dtos.SelectResponse, error) {
	var res *[]dtos.SelectResponse = &[]dtos.SelectResponse{}

	rows, err := db.Query("Select * from task where username=$1;", username)
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить задачи: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var row *dtos.SelectResponse = &dtos.SelectResponse{}
		if err := rows.Scan(&row.Id, &row.Username, &row.Title, &row.Description, &row.Due_date, &row.Done); err != nil {
			return nil, fmt.Errorf("Не удалось просканировать строку: %w", err)
		}
		*res = append(*res, *row)
	}

	return res, nil
}

func SelectCurrentTask(username string, id int) (*dtos.SelectResponse, error) {
	var res *dtos.SelectResponse = &dtos.SelectResponse{}
	err := db.QueryRow("Select * from task where username=$1 and id=$2;", username, id).Scan(&res.Id, &res.Username, &res.Title, &res.Description, &res.Due_date, &res.Done)
	if err != nil {
		return nil, fmt.Errorf("Не удалось получить задачу: %w", err)
	}

	return res, nil
}

func UpdateTask(username string, id int, req *dtos.UpdateTaskRequest) (*dtos.SelectResponse, error) {
	query := "UPDATE task SET "
	var updatedFields string
	params := []interface{}{}
	paramCount := 1

	if req.Title != nil {
		updatedFields += fmt.Sprintf("title=$%d, ", paramCount)
		params = append(params, *req.Title)
		paramCount++
	}

	if req.Description != nil {
		updatedFields += fmt.Sprintf("description=$%d, ", paramCount)
		params = append(params, *req.Description)
		paramCount++
	}

	if req.Due_date != nil {
		updatedFields += fmt.Sprintf("due_date=$%d, ", paramCount)
		params = append(params, *req.Due_date)
		paramCount++
	}

	if req.Done != nil {
		updatedFields += fmt.Sprintf("done=$%d, ", paramCount)
		params = append(params, *req.Done)
		paramCount++
	}

	if len(updatedFields) == 0 {
		return nil, fmt.Errorf("Нечего обновлять!")
	}

	updatedFields = updatedFields[:len(updatedFields)-2]

	query += fmt.Sprintf("%s WHERE username=$%d and id=$%d RETURNING *;", updatedFields, paramCount, paramCount+1)
	params = append(params, username)
	params = append(params, id)

	var res *dtos.SelectResponse = &dtos.SelectResponse{}

	err := db.QueryRow(query, params...).Scan(&res.Id, &res.Username, &res.Title, &res.Description, &res.Due_date, &res.Done)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при исполнении запроса Update: %w", err)
	}

	return res, nil
}

func DeleteTask(username string, id int) error {
	_, err := db.Exec("DELETE FROM task WHERE username=$1 and id=$2;", username, id)
	if err != nil {
		return fmt.Errorf("Ошибка при удалении: %w", err)
	}
	return nil
}
