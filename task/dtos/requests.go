package dtos

import "time"

type InsertRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Due_date    time.Time `json:"due_date"`
}

type UpdateTaskRequest struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	Due_date    *time.Time `json:"due_date,omitempty"`
	Done        *bool      `json:"done,omitempty"`
}
