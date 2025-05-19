package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ChileKasoka/construction-app/service"
	"github.com/go-chi/chi/v5"
)

type RolePermissionController struct {
	Service *service.RolePermissionService
}

func NewRolePermissionController(service *service.RolePermissionService) *RolePermissionController {
	return &RolePermissionController{Service: service}
}

type AssignPayload struct {
	PermissionIDs []int `json:"permission_ids"`
}

func (c *RolePermissionController) AssignPermission(w http.ResponseWriter, r *http.Request) {
	roleID, err := strconv.Atoi(chi.URLParam(r, "roleID"))
	if err != nil {
		http.Error(w, "invalid role id", http.StatusBadRequest)
		return
	}

	var payload AssignPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if len(payload.PermissionIDs) == 0 {
		http.Error(w, "no permissions provided", http.StatusBadRequest)
		return
	}

	err = c.Service.AssignPermissions(roleID, payload.PermissionIDs)
	if err != nil {
		http.Error(w, "could not assign permissions", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *RolePermissionController) RevokePermission(w http.ResponseWriter, r *http.Request) {
	roleID, _ := strconv.Atoi(chi.URLParam(r, "roleID"))
	permissionID, _ := strconv.Atoi(chi.URLParam(r, "permissionID"))

	err := c.Service.RevokePermission(roleID, permissionID)
	if err != nil {
		http.Error(w, "Could not revoke permission", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *RolePermissionController) ListPermissions(w http.ResponseWriter, r *http.Request) {
	roleID, _ := strconv.Atoi(chi.URLParam(r, "roleID"))

	permissions, err := c.Service.ListPermissions(roleID)
	if err != nil {
		http.Error(w, "Could not fetch permissions", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(permissions)
}

func (c *RolePermissionController) ListAllRolePermissions(w http.ResponseWriter, r *http.Request) {
	result, err := c.Service.GetAllRolePermissions()
	if err != nil {
		http.Error(w, "failed to load role-permission mappings", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
