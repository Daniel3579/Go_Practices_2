package network

import (
	"encoding/json"
	"log"
	"net/http"

	"prc-13/services/tasks/core/jobs"

	"github.com/google/uuid"
)

type JobPublisher interface {
	PublishJob(queueName string, job *jobs.TaskJob) error
}

type JobHandler struct {
	publisher JobPublisher
	queueName string
}

func NewJobHandler(publisher JobPublisher, queueName string) *JobHandler {
	return &JobHandler{
		publisher: publisher,
		queueName: queueName,
	}
}

type postJobRequest struct {
	TaskID string `json:"task_id"`
}

type postJobResponse struct {
	Status string `json:"status"`
	TaskID string `json:"task_id"`
}

func (h *JobHandler) ProcessTask(w http.ResponseWriter, r *http.Request) {
	var req postJobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if req.TaskID == "" {
		http.Error(w, "task_id is required", http.StatusBadRequest)
		return
	}

	job := &jobs.TaskJob{
		Job:       "process_task",
		TaskID:    req.TaskID,
		Attempt:   1,
		MessageID: uuid.New().String(),
	}

	if err := h.publisher.PublishJob(h.queueName, job); err != nil {
		log.Printf("failed to publish job: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(postJobResponse{
		Status: "accepted",
		TaskID: req.TaskID,
	})
}
