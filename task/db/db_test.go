package db

import (
	"fmt"
	"log"
	"os"
	"separation/task/dtos"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// ——————————————————————————————————————————————————————————————————————————————

func createTable() error {
	_, err := db.Exec(`
		CREATE TABLE task (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			title VARCHAR(50) NOT NULL,
			description VARCHAR(255),
			due_date DATE,
			done BOOLEAN DEFAULT FALSE
		);
	`)

	if err != nil {
		return fmt.Errorf("Ошибка при создании таблицы: %w", err)
	}

	return nil
}

func dropTable() error {
	_, err := db.Exec(`Drop table task;`)

	if err != nil {
		return fmt.Errorf("Ошибка при удалении таблицы: %w", err)
	}

	return nil
}

// ——————————————————————————————————————————————————————————————————————————————

func initTest() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env: %w", err)
	}

	err = ConnectDB("TEST_DATABASE_URL")
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func terminationTest() {
	dropTable()
	db.Close()
}

// ——————————————————————————————————————————————————————————————————————————————

func TestMain(m *testing.M) {
	initTest()
	code := m.Run()
	terminationTest()
	os.Exit(code)
}

// ——————————————————————————————————————————————————————————————————————————————

func TestInsert(t *testing.T) {
	var cases = []struct {
		username string
		req      dtos.InsertRequest
		want     dtos.SelectResponse
	}{
		{
			"email.com",
			dtos.InsertRequest{
				Title:       "Note one",
				Description: "Description nice",
				Due_date:    time.Date(2026, 2, 28, 0, 0, 0, 0, time.UTC),
			},
			dtos.SelectResponse{
				Id:          1,
				Username:    "email.com",
				Title:       "Note one",
				Description: "Description nice",
				Due_date:    time.Date(2026, 2, 28, 0, 0, 0, 0, time.UTC),
				Done:        false,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.req.Title, func(t *testing.T) {
			err := InsertIntoTask(c.username, c.req)
			assert.NoError(t, err)
		})
	}
}

func TestSelectAll(t *testing.T) {
	var cases = []struct {
		username string
		req      []dtos.InsertRequest
		want     []dtos.SelectResponse
	}{
		{
			"email.com",
			[]dtos.InsertRequest{
				{Title: "Note one", Description: "Description nice", Due_date: time.Date(2026, 2, 28, 0, 0, 0, 0, time.UTC)},
				{Title: "Note two", Description: "Description twice", Due_date: time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC)},
			},
			[]dtos.SelectResponse{
				{Id: 1, Username: "email.com", Title: "Note one", Description: "Description nice", Due_date: time.Date(2026, 2, 28, 0, 0, 0, 0, time.UTC), Done: false},
				{Id: 2, Username: "email.com", Title: "Note two", Description: "Description twice", Due_date: time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC), Done: false},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.username, func(t *testing.T) {
			err := InsertIntoTask(c.username, c.req[0])
			err = InsertIntoTask(c.username, c.req[1])

			res, err := SelectAllTasks(c.username)
			assert.NoError(t, err)

			for i := range 2 {
				assert.Equal(t, c.want[i].Id, res[i].Id)
				assert.Equal(t, c.want[i].Username, res[i].Username)
				assert.Equal(t, c.want[i].Title, res[i].Title)
				assert.Equal(t, c.want[i].Description, res[i].Description)
				assert.Equal(t, c.want[i].Done, res[i].Done)
				assert.True(t, c.want[i].Due_date.Equal(res[i].Due_date))
			}
		})
	}
}

func ptr[T any](value T) *T {
	return &value
}

func TestUpdateTask(t *testing.T) {
	var cases = []struct {
		username string
		id       int
		req      dtos.UpdateTaskRequest
		want     dtos.SelectResponse
	}{
		{
			"email.com",
			1,
			dtos.UpdateTaskRequest{
				Title:       ptr("123"),
				Description: nil,
				Due_date:    nil,
				Done:        nil,
			},
			dtos.SelectResponse{Id: 1, Username: "email.com", Title: "123", Description: "Description nice", Due_date: time.Date(2026, 2, 28, 0, 0, 0, 0, time.UTC), Done: false},
		},
	}

	for _, c := range cases {
		t.Run(c.username, func(t *testing.T) {
			err := InsertIntoTask("email.com", dtos.InsertRequest{Title: "Note one", Description: "Description nice", Due_date: time.Date(2026, 2, 28, 0, 0, 0, 0, time.UTC)})
			err = UpdateTask(c.username, c.id, c.req)
			if err != nil {
				fmt.Println(err.Error())
			}
		})
	}
}
