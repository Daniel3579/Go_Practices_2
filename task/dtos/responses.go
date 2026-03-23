package dtos

import "time"

type SelectResponse struct {
	Id          int       `json:"id"`
	Username    string    `json:"username"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Due_date    time.Time `json:"due_date"`
	Done        bool      `json:"done"`
}
