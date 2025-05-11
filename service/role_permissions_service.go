package service

import (
	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
)

type RolePermissionService struct {
	Repo *repository.RolePermissionRepo
}

func (s *RolePermissionService) AssignPermission(roleID, permissionID int) error {
	return s.Repo.Create(roleID, permissionID)
}

func (s *RolePermissionService) RevokePermission(roleID, permissionID int) error {
	return s.Repo.Delete(roleID, permissionID)
}

func (s *RolePermissionService) ListPermissions(roleID int) ([]model.Permission, error) {
	return s.Repo.GetByRoleID(roleID)
}
