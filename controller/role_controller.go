package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/service"
	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	Service *service.RoleService
}

func NewRoleController(service *service.RoleService) *RoleController {
	return &RoleController{Service: service}
}

func (c *RoleController) Create(w http.ResponseWriter, r *http.Request) {
	var req model.Role
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

}

func (c *RoleController) GetAll(w http.ResponseWriter, r *http.Request) {
	roles, err := c.Service.GetAll()
	if err != nil {
		http.Error(w, "failed to get roles", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(roles)
}

func (c *RoleController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	role, err := c.Service.GetByID(id)
	if err != nil {
		http.Error(w, "role not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(role)
}

func (c *RoleController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	var role model.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	role.ID = id
	err = c.Service.Update(&role)
	if err != nil {
		http.Error(w, "failed to update role", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *RoleController) FindByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	role, err := c.Service.FindByName(name)
	if err != nil {
		http.Error(w, "role not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(role)
}

func (c *RoleController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := c.Service.Delete(id)
	if err != nil {
		http.Error(w, "failed to delete role", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
