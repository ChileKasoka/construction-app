package service

import (
	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
)

type PermissionService struct {
	Repo *repository.PermissionRepository
}

func NewPermissionService(repo *repository.PermissionRepository) *PermissionService {
	return &PermissionService{Repo: repo}
}

func (s *PermissionService) Create(permission *model.Permission) error {
	return s.Repo.Create(permission)
}

func (s *PermissionService) GetAll() ([]model.Permission, error) {
	return s.Repo.GetAll()
}

func (s *PermissionService) GetByID(id int64) (*model.Permission, error) {
	return s.Repo.GetByID(id)
}

func (s *PermissionService) GetUnassignedByRoleID(roleID int) ([]model.Permission, error) {
	return s.Repo.GetUnassignedByRoleID(roleID)
}

func (s *PermissionService) Update(p *model.Permission) error {
	return s.Repo.Update(p)
}

func (s *PermissionService) Delete(id int64) error {
	return s.Repo.Delete(id)
}
