package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var (
	tasks = map[string]Task{
		"t_001": {ID: "t_001", Title: "Изучить REST", Description: "Прочитать про REST API", Done: false},
		"t_002": {ID: "t_002", Title: "Изучить GraphQL", Description: "Разобраться со схемами", Done: true},
	}
	mu sync.RWMutex
)

func main() {
	http.HandleFunc("/v1/tasks", tasksHandler)
	http.HandleFunc("/v1/tasks/", taskHandler) // для /v1/tasks/{id}
	log.Println("REST server on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		defer mu.RUnlock()
		list := make([]Task, 0, len(tasks))
		for _, t := range tasks {
			list = append(list, t)
		}
		json.NewEncoder(w).Encode(list)
	case http.MethodPost:
		var newTask Task
		if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		newTask.ID = uuid.New().String()
		newTask.Done = false
		mu.Lock()
		tasks[newTask.ID] = newTask
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/v1/tasks/")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		task, ok := tasks[id]
		mu.RUnlock()
		if !ok {
			http.Error(w, "task not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(task)
	case http.MethodPatch:
		var updates map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		task, ok := tasks[id]
		if !ok {
			http.Error(w, "task not found", http.StatusNotFound)
			return
		}
		if doneVal, ok := updates["done"]; ok {
			if done, ok := doneVal.(bool); ok {
				task.Done = done
			}
		}
		tasks[id] = task
		json.NewEncoder(w).Encode(task)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
