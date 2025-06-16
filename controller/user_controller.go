package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/service"
	"github.com/go-chi/chi/v5"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	ID          int                `json:"id"`
	AccessToken string             `json:"access_token"`
	Role        string             `json:"role"`
	User        model.User         `json:"user"`
	RoleID      int                `json:"role_id"`
	Permissions []model.Permission `json:"permissions"`
}

type RegisterResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	RoleID    int    `json:"role_id"`
	RoleName  string `json:"role_name"`
	RoleDesc  string `json:"role_desc"`
	CreatedAt string `json:"created_at"`
}

type UserController struct {
	Service               *service.UserService
	RolePermissionService *service.RolePermissionService
	JWTSecret             string
}

func NewUserController(service *service.UserService, rolePermissionService *service.RolePermissionService) *UserController {
	return &UserController{
		Service:               service,
		RolePermissionService: rolePermissionService,
	}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// role, err :=

	err := c.Service.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	accessToken, _, role, user, roleID, err := c.Service.Authenticate(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	permissions, err := c.RolePermissionService.ListPermissions(roleID)
	if err != nil {
		http.Error(w, "failed to fetch permissions", http.StatusInternalServerError)
		return
	}

	res := LoginResponse{
		ID:          user.ID,
		AccessToken: accessToken,
		Role:        role,
		User:        user,
		Permissions: permissions,
		RoleID:      roleID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, _ := c.Service.GetAll()
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetAllCount(w http.ResponseWriter, r *http.Request) {
	count, err := c.Service.GetAllCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(count)
}

func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := c.Service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user.ID = id // Ensure path param is used
	err = c.Service.Update(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := c.Service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
