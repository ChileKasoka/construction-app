package service

import (
	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
)

type ProjectService struct {
	Repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{Repo: repo}
}

func (s *ProjectService) GetAll() ([]model.Project, error) {
	return s.Repo.GetAll()
}

func (s *ProjectService) GetByID(id int) (*model.Project, error) {
	return s.Repo.GetByID(id)
}

func (s *ProjectService) Create(p *model.Project) (*model.Project, error) {
	return s.Repo.Create(p)
}

func (s *ProjectService) Update(id int, p *model.Project) (*model.Project, error) {
	return s.Repo.Update(id, p)
}

func (s *ProjectService) Delete(id int) error {
	return s.Repo.Delete(id)
}
