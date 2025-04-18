package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/service"
	"github.com/go-chi/chi/v5"
)

type ProjectController struct {
	Service *service.ProjectService
}

func (c *ProjectController) GetAll(w http.ResponseWriter, r *http.Request) {
	projects, _ := c.Service.GetAll()
	json.NewEncoder(w).Encode(projects)
}

func (c *ProjectController) Create(w http.ResponseWriter, r *http.Request) {
	var p model.Project
	json.NewDecoder(r.Body).Decode(&p)
	created, err := c.Service.Create(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(created)
}

func (c *ProjectController) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var p model.Project
	json.NewDecoder(r.Body).Decode(&p)
	updated, err := c.Service.Update(id, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (c *ProjectController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := c.Service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
