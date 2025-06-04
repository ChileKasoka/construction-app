package service

import (
	"fmt"

	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
)

type UserTaskService struct {
	userTaskRepo *repository.UserTaskRepository
	taskRepo     *repository.TaskRepository
}

func NewUserTaskService(repo *repository.UserTaskRepository, repoTask *repository.TaskRepository) *UserTaskService {
	return &UserTaskService{
		userTaskRepo: repo,
		taskRepo:     repoTask,
	}
}

func (s *UserTaskService) Create(userTask *model.UserTask) error {
	task, err := s.taskRepo.GetByID(userTask.TaskID)
	if err != nil {
		return fmt.Errorf("failed to fetch task: %w", err)
	}
	if task == nil {
		return fmt.Errorf("task with ID %d does not exist", userTask.TaskID)
	}

	err = s.userTaskRepo.Create(userTask)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserTaskService) AssignUsersToTask(taskID int, userIDs []int) error {
	// Optional: validate userIDs or check if task exists
	if len(userIDs) == 0 {
		return nil
	}
	return s.userTaskRepo.CreateMany(taskID, userIDs)
}

func (s *UserTaskService) GetAll() ([]model.UserTask, error) {
	userTasks, err := s.userTaskRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return userTasks, nil
}

func (s *UserTaskService) GetByUserID(userID int) ([]model.UserTask, error) {
	userTasks, err := s.userTaskRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return userTasks, nil
}

func (s *UserTaskService) UnassignUser(userID, taskID int) error {
	err := s.userTaskRepo.UnassignUser(userID, taskID)
	if err != nil {
		return err
	}
	return nil
}
