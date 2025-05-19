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
	roleRepo := repository.NewRoleRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)
	rolePermissionRepo := repository.NewRolePerissionRepo(db)

	// Set up services
	userService := service.NewUserService(userRepo, roleRepo)
	projectService := service.NewProjectService(projectRepo)
	roleService := service.NewRoleService(roleRepo)
	permissionService := service.NewPermissionService(permissionRepo)
	rolePermissionService := service.NewRolePermissionService(rolePermissionRepo)

	// Set up controllers
	userController := controller.NewUserController(userService)
	projectController := controller.NewProjectController(projectService)
	roleController := controller.NewRoleController(roleService)
	permissionController := controller.NewPermissionController(permissionService)
	rolePermissionController := controller.NewRolePermissionController(rolePermissionService)

	// Set up router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mw.CorsMiddleware)

	r.Route("/projects", func(r chi.Router) {
		r.With(mw.RoleMiddleware("admin")).Get("/", projectController.GetAll)
		r.Post("/", projectController.Create)
		r.Get("/{id}", projectController.GetByID)
		r.Put("/{id}", projectController.Update)
		r.Delete("/{id}", projectController.Delete)
	})

	r.Route("/roles", func(r chi.Router) {
		r.Get("/", roleController.GetAll)
		r.Post("/", roleController.Create)
		r.With(mw.RoleMiddleware("admin")).Get("/{id}", roleController.GetByID)
		r.Put("/{id}", roleController.Update)
		r.Delete("/{id}", roleController.Delete)
		r.Get("/name/{name}", roleController.FindByName)
	})

	r.Route("/permissions", func(r chi.Router) {
		r.Post("/", permissionController.Create)
		r.Get("/", permissionController.GetAll)
	})

	r.Route("/role-permissions", func(r chi.Router) {
		r.Post("/{roleID}", rolePermissionController.AssignPermission)
		r.Get("/", rolePermissionController.ListAllRolePermissions)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userController.Create)
		r.With(mw.RoleMiddleware("admin")).Get("/{id}", userController.GetByID)
		r.Put("/{id}", userController.Update)
		r.Delete("/{id}", userController.Delete)
	})
	r.Post("/login", userController.Login)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
