package service

import (
	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetByID(id int) (*model.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) Update(id int, data map[string]interface{}) (*model.User, error) {
	return s.Repo.Update(id, data)
}

func (s *UserService) Delete(id int) error {
	return s.Repo.Delete(id)
}
