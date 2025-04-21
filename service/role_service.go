package service

import (
	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
)

type RoleService struct {
	Repo *repository.RoleRepository
}

func NewRoleService(repo *repository.RoleRepository) *RoleService {
	return &RoleService{Repo: repo}
}

func (rs *RoleService) Create(role *model.Role) error {
	return rs.Repo.Create(role)
}

func (rs *RoleService) GetAll() ([]model.Role, error) {
	return rs.Repo.GetAll()
}

func (rs *RoleService) GetByID(id int) (*model.Role, error) {
	return rs.Repo.GetByID(id)
}

func (rs *RoleService) Update(role *model.Role) error {
	return rs.Repo.Update(role)
}

func (rs *RoleService) FindByName(name string) (*model.Role, error) {
	return rs.Repo.FindByName(name)
}

func (rs *RoleService) Delete(id int) error {
	return rs.Repo.Delete(id)
}
