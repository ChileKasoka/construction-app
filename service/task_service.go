package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
)

type TaskService struct {
	TaskRepo *repository.TaskRepository
}

func NewTaskService(service *repository.TaskRepository) *TaskService {
	return &TaskService{
		TaskRepo: service,
	}
}

func (s *TaskService) Create(req model.Task) error {
	// Basic validation (optional, expand as needed)
	if req.Title == "" {
		return errors.New("task title is required")
	}

	// Set timestamps
	now := time.Now()
	req.CreatedAt = now
	req.UpdatedAt = now

	// Pass to repository
	err := s.TaskRepo.Create(&req)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (s *TaskService) GetByID(id int) (model.Task, error) {
	task, err := s.TaskRepo.GetByID(id)
	if err != nil {
		return model.Task{}, fmt.Errorf("failed to fetch task: %w", err)
	}
	if task == nil {
		return model.Task{}, fmt.Errorf("task with ID %d not found", id)
	}

	return *task, nil
}

func (s *TaskService) GetAll() ([]model.Task, error) {
	tasks, err := s.TaskRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}
	return tasks, nil
}

func (s *TaskService) GetAllCount() (int, error) {
	count, err := s.TaskRepo.GetAllCount()
	if err != nil {
		return 0, fmt.Errorf("failed to get task count: %w", err)
	}
	return count, nil
}

func (s *TaskService) Update(id int, req *model.Task) error {
	// Basic validation (optional, expand as needed)
	if req.Title == "" {
		return errors.New("task title is required")
	}

	// Set updated timestamp
	req.UpdatedAt = time.Now()

	// Pass to repository
	err := s.TaskRepo.Update(id, req)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}

func (s *TaskService) Delete(id int) error {
	return s.TaskRepo.Delete(id)
}
