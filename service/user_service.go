package service

import (
	"errors"
	"math/rand"

	"github.com/ChileKasoka/construction-app/middleware/auth"
	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
	// "golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo     *repository.UserRepository
	RoleRepo *repository.RoleRepository
}

func NewUserService(uRepo *repository.UserRepository, rRepo *repository.RoleRepository) *UserService {
	return &UserService{
		Repo:     uRepo,
		RoleRepo: rRepo,
	}
}

func generateRandomPassword(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, n)
	for i := range password {
		password[i] = letters[rand.Intn(len(letters))]
	}
	return string(password)
}

func (s *UserService) Create(req model.RegisterRequest) error {
	// TODO: add validation or email uniqueness check if needed

	role, err := s.RoleRepo.GetByID(req.RoleID)
	if err != nil {
		return errors.New("role not found")
	}

	password := req.Password
	if req.Password == "" {
		password = generateRandomPassword(10)
	}
	user := &model.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: password,
		RoleID:   role.ID,
	}
	return s.Repo.Create(user)
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetByID(id int) (*model.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) Authenticate(email, password string) (string, string, string, error) {
	// Fetch user by email
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", "", "", errors.New("invalid email or password")
	}

	// Compare hashed password
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// if err != nil {
	// 	return "", "", errors.New("invalid email or password")
	// }

	// Generate JWT token
	token, err := auth.CreateJWT(user.ID, user.Role.Name)
	if err != nil {
		return "", "", "", errors.New("failed to generate token")
	}

	return token, user.Role.Name, user.Name, nil
}

func (s *UserService) Update(user model.User) error {
	return s.Repo.Update(user)
}

func (s *UserService) Delete(id int) error {
	return s.Repo.Delete(id)
}
