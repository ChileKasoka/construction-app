package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/service"
	"github.com/go-chi/chi/v5"
)

type TaskController struct {
	Service *service.TaskService
}

func NewTaskController(service *service.TaskService) *TaskController {
	return &TaskController{Service: service}
}

func (c *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Call service to create task
	if err := c.Service.Create(task); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create task: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with 201 Created
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Task created successfully"}`))
}

func (c *TaskController) GetByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL (assuming you're using chi)
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Fetch task from service
	task, err := c.Service.GetByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch task: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with task JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (c *TaskController) GetAll(w http.ResponseWriter, r *http.Request) {
	// Fetch tasks from the service layer
	tasks, err := c.Service.GetAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch tasks: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (c *TaskController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := c.Service.Update(id, &task); err != nil {
		http.Error(w, fmt.Sprintf("Failed to update task: %v", err), http.StatusInternalServerError)
		return
	}

	// Set the task ID so it's returned in the response
	task.ID = id

	// Return the updated task as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (c *TaskController) Delete(w http.ResponseWriter, r *http.Request) {
	// Get task ID from URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Call service to delete the task
	if err := c.Service.Delete(id); err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete task: %v", err), http.StatusInternalServerError)
		return
	}

	// Success: 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
