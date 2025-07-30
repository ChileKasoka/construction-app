package service

import (
	"fmt"

	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
)

type UserProjectService struct {
	userProjectRepo *repository.UserProjectRepository
}

func NewUserProjectService(repo *repository.UserProjectRepository) *UserProjectService {
	return &UserProjectService{
		userProjectRepo: repo,
	}
}

func (s *UserProjectService) Create(userProject *model.UserProject) error {
	existingUserProject, err := s.userProjectRepo.GetByProjectID(userProject.ProjectID)
	if err != nil {
		return fmt.Errorf("failed to fetch user project: %w", err)
	}
	if existingUserProject != nil {
		return fmt.Errorf("user project with ID %d already exists", userProject.ProjectID)
	}
	return s.userProjectRepo.Create(userProject)
}

func (s *UserProjectService) CreateMany(projectID int, userIDs []int) error {
	return s.userProjectRepo.CreateMany(projectID, userIDs)
}

func (s *UserProjectService) GetAll() ([]model.User, error) {
	return s.userProjectRepo.GetAll()
}

func (s *UserProjectService) GetByProjectID(projectID int) ([]model.ProjectUsers, error) {
	return s.userProjectRepo.GetByProjectID(projectID)
}

func (s *UserProjectService) GetByUserID(userID int) ([]model.Project, error) {
	return s.userProjectRepo.GetByUserID(userID)
}
