package service

import (
	"errors"

	"github.com/ChileKasoka/construction-app/middleware/auth"
	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
	// "golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
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

func (s *UserService) Authenticate(email, password string) (string, string, error) {
	// Fetch user by email
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	// Compare hashed password
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// if err != nil {
	// 	return "", "", errors.New("invalid email or password")
	// }

	// Generate JWT token
	token, err := auth.CreateJWT(user.ID, user.Role.Name)
	if err != nil {
		return "", "", errors.New("failed to generate token")
	}

	return token, user.Role.Name, nil
}

func (s *UserService) Update(user model.User) error {
	return s.Repo.Update(user)
}

func (s *UserService) Delete(id int) error {
	return s.Repo.Delete(id)
}
