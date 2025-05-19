package service

import (
	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
)

type RolePermissionService struct {
	Repo *repository.RolePermissionRepo
}

func NewRolePermissionService(repo *repository.RolePermissionRepo) *RolePermissionService {
	return &RolePermissionService{Repo: repo}
}

func (s *RolePermissionService) AssignPermissions(roleID int, permissionIDs []int) error {
	return s.Repo.Create(roleID, permissionIDs)
}

func (s *RolePermissionService) RevokePermission(roleID, permissionID int) error {
	return s.Repo.Delete(roleID, permissionID)
}

func (s *RolePermissionService) ListPermissions(roleID int) ([]model.Permission, error) {
	return s.Repo.GetByRoleID(roleID)
}

func (s *RolePermissionService) GetAllRolePermissions() ([]repository.RolePermissionDisplay, error) {
	return s.Repo.GetAllRolePermissions()
}
