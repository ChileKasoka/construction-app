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

type AssignPayload struct {
	PermissionID int `json:"permission_id"`
}

func (c *RolePermissionController) AssignPermission(w http.ResponseWriter, r *http.Request) {
	roleID, _ := strconv.Atoi(chi.URLParam(r, "roleID"))

	var payload AssignPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := c.Service.AssignPermission(roleID, payload.PermissionID)
	if err != nil {
		http.Error(w, "could not assign permission", http.StatusInternalServerError)
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
