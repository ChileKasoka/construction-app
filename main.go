package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ChileKasoka/construction-app/controller"
	mw "github.com/ChileKasoka/construction-app/middleware"
	"github.com/ChileKasoka/construction-app/repository"
	"github.com/ChileKasoka/construction-app/service"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Project setup
	projectRepo := &repository.ProjectRepository{}
	projectService := &service.ProjectService{Repo: projectRepo}
	projectController := &controller.ProjectController{Service: projectService}

	r.Get("/projects", projectController.GetAll)
	r.Get("/projects/{id}", projectController.GetByID)

	r.Group(func(r chi.Router) {
		r.Use(mw.RoleMiddleware("admin"))
		r.Post("/projects", projectController.Create)
		r.Put("/projects/{id}", projectController.Update)
		r.Delete("/projects/{id}", projectController.Delete)
	})

	// User setup
	userRepo := &repository.UserRepository{}
	userService := &service.UserService{Repo: userRepo}
	userController := &controller.UserController{Service: userService}

	r.Group(func(r chi.Router) {
		r.Use(mw.RoleMiddleware("admin"))
		r.Get("/users", userController.GetAll)
		r.Get("/users/{id}", userController.GetByID)
		r.Put("/users/{id}", userController.Update)
		r.Delete("/users/{id}", userController.Delete)
	})

	http.ListenAndServe(":8080", r)
}
