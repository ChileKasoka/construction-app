package service

import (
	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) Create(user *model.User) error {
	// TODO: add validation or email uniqueness check if needed
	return s.Repo.Create(user)
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetByID(id int) (*model.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) Update(user model.User) error {
	return s.Repo.Update(user)
}

func (s *UserService) Delete(id int) error {
	return s.Repo.Delete(id)
}
