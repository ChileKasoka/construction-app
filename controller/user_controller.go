package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ChileKasoka/construction-app/service"
	"github.com/go-chi/chi/v5"
)

type UserController struct {
	Service *service.UserService
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, _ := c.Service.GetAll()
	json.NewEncoder(w).Encode(users)
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
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var user map[string]interface{}
	json.NewDecoder(r.Body).Decode(&user)
	updated, err := c.Service.Update(id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updated)
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
