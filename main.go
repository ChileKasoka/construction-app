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

	// cfg := config.LoadConfig()

	// Set up repositories
	rolePermRepo := repository.NewRolePerissionRepo(db)

	userRepo := repository.NewUserRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)
	rolePermissionRepo := repository.NewRolePerissionRepo(db)
	taskRepo := repository.NewTaskRepository(db)
	userTaskRepo := repository.NewUserTaskRepository(db)
	userProjectRepo := repository.NewUserProjectRepository(db)

	// Set up services
	userService := service.NewUserService(userRepo, roleRepo)
	projectService := service.NewProjectService(projectRepo)
	roleService := service.NewRoleService(roleRepo)
	permissionService := service.NewPermissionService(permissionRepo)
	rolePermissionService := service.NewRolePermissionService(rolePermissionRepo)
	taskService := service.NewTaskService(taskRepo)
	userTaskService := service.NewUserTaskService(userTaskRepo, taskRepo)
	userProjectService := service.NewUserProjectService(userProjectRepo)

	// Set up controllers
	userController := controller.NewUserController(userService, rolePermissionService)
	projectController := controller.NewProjectController(projectService)
	roleController := controller.NewRoleController(roleService)
	permissionController := controller.NewPermissionController(permissionService)
	rolePermissionController := controller.NewRolePermissionController(rolePermissionService)
	taskController := controller.NewTaskController(taskService)
	userTaskController := controller.NewUserTaskController(userTaskService)
	userProjectController := controller.NewUserProjectController(userProjectService)

	// Set up router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mw.CorsMiddleware)

	r.Route("/projects", func(r chi.Router) {
		r.Get("/", projectController.GetAll)
		r.Post("/", projectController.Create)
		r.Get("/{id}", projectController.GetByID)
		r.Put("/{id}", projectController.Update)
		r.Delete("/{id}", projectController.Delete)
	})

	r.Route("/roles", func(r chi.Router) {
		r.Use(mw.RoleMiddleware(rolePermRepo))
		r.Get("/", roleController.GetAll)
		r.Post("/", roleController.Create)
		r.Get("/{id}", roleController.GetByID)
		r.Put("/{id}", roleController.Update)
		r.Delete("/{id}", roleController.Delete)
		r.Get("/name/{name}", roleController.FindByName)
	})

	r.Route("/user-tasks", func(r chi.Router) {
		r.Post("/", userTaskController.Create)
		r.Get("/", userTaskController.GetAll)
		r.Get("/{id}", userTaskController.GetByUserID)
		r.Delete("/{user_id}", userTaskController.UnassignUserFromTask)
	})

	r.Route("/user-projects", func(r chi.Router) {
		r.Post("/{projectID}/assign-users", userProjectController.Create)
		r.Post("/{projectID}/many", userProjectController.CreateMany)
		r.Get("/", userProjectController.GetAll)
		// r.Get("/{user_id}/projects", userProjectController.GetByProjectID)
		r.Get("/{user_id}/projects", userProjectController.GetByUserID)
	})

	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", taskController.Create)
		r.Get("/", taskController.GetAll)
		r.Get("/{id}", taskController.GetByID)
		r.Put("/{id}", taskController.Update)
		r.Delete("/{id}", taskController.Delete)
		r.Post("/{id}/assign-users", userTaskController.AssignUsersToTaskHandler)
	})

	r.Route("/permissions", func(r chi.Router) {
		r.Use(mw.RoleMiddleware(rolePermRepo))
		r.Post("/", permissionController.Create)
		r.Get("/", permissionController.GetAll)
		r.Get("/{id}", permissionController.GetByID)
	})

	r.Route("/role-permissions", func(r chi.Router) {
		r.Use(mw.RoleMiddleware(rolePermRepo))
		r.Post("/{roleID}", rolePermissionController.AssignPermission)
		r.Get("/", rolePermissionController.ListAllRolePermissions)
		r.Get("/{id}/permissions", rolePermissionController.ListPermissions)
		r.Get("/{id}", rolePermissionController.GetByUserID)
	})

	r.Route("/users", func(r chi.Router) {
		// r.Use(mw.RoleMiddleware(rolePermRepo))
		r.Get("/count", userController.GetAllCount)
		r.Post("/", userController.Create)
		r.Get("/", userController.GetAll)
		r.Get("/{id}", userController.GetByID)
		r.Put("/{id}", userController.Update)
		r.Delete("/{id}", userController.Delete)
	})
	r.Post("/login", userController.Login)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
