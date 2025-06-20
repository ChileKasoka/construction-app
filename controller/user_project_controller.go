package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/service"
	"github.com/go-chi/chi/v5"
)

type UserProjectController struct {
	Service *service.UserProjectService
}

func NewUserProjectController(service *service.UserProjectService) *UserProjectController {
	return &UserProjectController{Service: service}
}

func (c *UserProjectController) Create(w http.ResponseWriter, r *http.Request) {
	var userProject model.UserProject
	if err := json.NewDecoder(r.Body).Decode(&userProject); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.Service.Create(&userProject); err != nil {
		http.Error(w, "Failed to create user project: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userProject)
}

func (c *UserProjectController) CreateMany(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ProjectID int   `json:"project_id"`
		UserIDs   []int `json:"user_ids"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(req.UserIDs) == 0 {
		http.Error(w, "No user IDs provided", http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateMany(req.ProjectID, req.UserIDs); err != nil {
		http.Error(w, "Failed to create user projects: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *UserProjectController) GetAll(w http.ResponseWriter, r *http.Request) {
	userProjects, err := c.Service.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch user projects: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userProjects)
}

func (c *UserProjectController) GetByProjectID(w http.ResponseWriter, r *http.Request) {
	projectIDStr := chi.URLParam(r, "projectID")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	users, err := c.Service.GetByProjectID(projectID)
	if err != nil {
		http.Error(w, "Failed to fetch users for project: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (c *UserProjectController) GetByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	projects, err := c.Service.GetByUserID(userID)
	if err != nil {
		http.Error(w, "Failed to fetch projects for user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
