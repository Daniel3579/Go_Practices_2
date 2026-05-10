package service

import (
	"fmt"
	"time"
)

type Task struct {
	ID          string
	Title       string
	Description string
}

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) CreateTask(title, description string) (*Task, error) {
	taskID := fmt.Sprintf("t_%d", time.Now().UnixNano())
	return &Task{
		ID:          taskID,
		Title:       title,
		Description: description,
	}, nil
}
