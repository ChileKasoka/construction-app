package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/service"
	"github.com/go-chi/chi/v5"
)

type UserTaskController struct {
	userTaskService *service.UserTaskService
}

type AssignTaskRequest struct {
	TaskID  int   `json:"task_id"`
	UserIDs []int `json:"user_ids"`
}

func NewUserTaskController(userTaskService *service.UserTaskService) *UserTaskController {
	return &UserTaskController{
		userTaskService: userTaskService,
	}
}

func (c *UserTaskController) Create(w http.ResponseWriter, r *http.Request) {
	var userTask model.UserTask
	if err := json.NewDecoder(r.Body).Decode(&userTask); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.userTaskService.Create(&userTask); err != nil {
		http.Error(w, "Failed to create user task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userTask)
}

func (c *UserTaskController) AssignUsersToTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var req AssignTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	req.TaskID = taskID

	err = c.userTaskService.AssignUsersToTask(req.TaskID, req.UserIDs)
	if err != nil {
		log.Printf("AssignUsersToTask error: %v\n", err)
		http.Error(w, "Failed to assign users to task: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // ✅ Proper status for POST creation
	json.NewEncoder(w).Encode(map[string]string{"message": "Users assigned successfully"})
}

func (c *UserTaskController) GetAll(w http.ResponseWriter, r *http.Request) {
	userTasks, err := c.userTaskService.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch user tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userTasks)
}

func (c *UserTaskController) GetByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convert userID from string to int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userTasks, err := c.userTaskService.GetByUserID(userID)
	if err != nil {
		http.Error(w, "Failed to fetch user tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userTasks)
}

func (h *UserTaskController) UnassignUserFromTask(w http.ResponseWriter, r *http.Request) {
	// Get userID and taskID from URL params
	userIDParam := chi.URLParam(r, "userID")
	taskIDParam := chi.URLParam(r, "taskID")

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	taskID, err := strconv.Atoi(taskIDParam)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Call the service method
	err = h.userTaskService.UnassignUser(userID, taskID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to unassign user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content is appropriate for a successful deletion
}
