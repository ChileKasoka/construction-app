package service

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/ChileKasoka/construction-app/middleware/auth"
	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
	email "github.com/ChileKasoka/construction-app/util"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserService) Create(req model.RegisterRequest) (string, error) {
	role, err := s.RoleRepo.GetByID(req.RoleID)
	if err != nil {
		return "", errors.New("role not found")
	}

	password := req.Password
	if password == "" {
		password = generateRandomPassword(10)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	user := &model.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		RoleID:   role.ID,
		// CreatedBy: req.CreatedBy, // Creator's email passed in request
	}

	err = s.Repo.Create(user)
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf(`
		<p>Hello %s,</p>
		<p>Your account has been created.</p>
		<p><strong>Email:</strong> %s</p>
		<p><strong>Password:</strong> %s</p>
		<p>You can now login and change your password.</p>
	`, req.Name, req.Email, password)

	// Send email to user and creator (non-blocking)
	go email.SendEmail(req.Email, "Your Account Credentials", message)
	// go email.SendEmail(req.CreatedBy, "User Created: "+req.Email, message)

	return password, nil
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetAllCount() (int, error) {
	count, err := s.Repo.GetAllCount()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *UserService) GetByID(id int) (*model.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) Authenticate(email, password string) (string, int, string, model.User, int, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", 0, "", model.User{}, 0, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Don't reveal which part is incorrect (email or password)
		return "", 0, "", model.User{}, 0, errors.New("invalid email or password")
	}

	role, err := s.RoleRepo.GetByID(user.RoleID)
	if err != nil {
		return "", 0, "", model.User{}, 0, errors.New("role not found")
	}

	token, err := auth.CreateJWT(user.ID, role.Name, role.ID)
	if err != nil {
		return "", 0, "", model.User{}, 0, errors.New("failed to generate token")
	}

	return token, user.ID, role.Name, *user, role.ID, nil
}

func (s *UserService) Update(user model.User) error {
	return s.Repo.Update(user)
}

func (s *UserService) Delete(id int) error {
	return s.Repo.Delete(id)
}
