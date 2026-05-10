package network

import (
	"encoding/json"
	"log"
	"net/http"

	"prc-13/services/tasks/core/service"
)

type TaskService interface {
	CreateTask(title, description string) (*service.Task, error)
}

type Publisher interface {
	PublishTaskCreated(taskID string)
}

type TaskHandler struct {
	service   TaskService
	publisher Publisher
}

func NewTaskHandler(service TaskService, publisher Publisher) *TaskHandler {
	return &TaskHandler{
		service:   service,
		publisher: publisher,
	}
}

type createTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type createTaskResponse struct {
	TaskID string `json:"task_id"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTask(req.Title, req.Description)
	if err != nil {
		log.Printf("failed to create task: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// Публикуем событие (best effort)
	h.publisher.PublishTaskCreated(task.ID)

	// Отвечаем клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createTaskResponse{TaskID: task.ID})
}
