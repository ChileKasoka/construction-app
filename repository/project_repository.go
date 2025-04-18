package repository

import (
	"errors"

	"github.com/ChileKasoka/construction-app/model"
)

type ProjectRepository struct{}

var projects = []model.Project{
	{ID: 1, Name: "Building A", Description: "Foundation and framing", Status: "In Progress"},
}

func (r *ProjectRepository) GetAll() ([]model.Project, error) {
	return projects, nil
}

func (r *ProjectRepository) GetByID(id int) (*model.Project, error) {
	for _, p := range projects {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, errors.New("project not found")
}

func (r *ProjectRepository) Create(p *model.Project) (*model.Project, error) {
	p.ID = len(projects) + 1
	projects = append(projects, *p)
	return p, nil
}

func (r *ProjectRepository) Update(id int, p *model.Project) (*model.Project, error) {
	for i, project := range projects {
		if project.ID == id {
			p.ID = id
			projects[i] = *p
			return p, nil
		}
	}
	return nil, errors.New("project not found")
}

func (r *ProjectRepository) Delete(id int) error {
	for i, project := range projects {
		if project.ID == id {
			projects = append(projects[:i], projects[i+1:]...)
			return nil
		}
	}
	return errors.New("project not found")
}
