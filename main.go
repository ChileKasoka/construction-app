package main

import (
	"log"
	"net/http"

	"github.com/ChileKasoka/construction-app/controller"
	"github.com/ChileKasoka/construction-app/db"
	mw "github.com/ChileKasoka/construction-app/middleware"
	"github.com/ChileKasoka/construction-app/repository"
	"github.com/ChileKasoka/construction-app/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db, err := db.ConnectDb()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Set up repositories
	userRepo := repository.NewUserRepository(db)
	projectRepo := repository.NewProjectRepository(db)

	// Set up services
	userService := service.NewUserService(userRepo)
	projectService := service.NewProjectService(projectRepo)

	// Set up controllers
	userController := controller.NewUserController(userService)
	projectController := controller.NewProjectController(projectService)

	// Set up router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/projects", func(r chi.Router) {
		r.With(mw.RoleMiddleware("admin")).Get("/", projectController.GetAll)
		r.Post("/", projectController.Create)
		r.Get("/{id}", projectController.GetByID)
		r.Put("/{id}", projectController.Update)
		r.Delete("/{id}", projectController.Delete)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", userController.GetByID)
		r.Put("/{id}", userController.Update)
		r.Delete("/{id}", userController.Delete)
	})
	r.Post("/login", userController.Login)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
