package repository

import (
	"database/sql"
	"fmt"

	"github.com/ChileKasoka/construction-app/model"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) Create(task *model.Task) error {
	query := `
	INSERT INTO tasks (title, description, status, start_date, end_date)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at, updated_at
	`
	err := r.DB.QueryRow(query, task.Title, task.Description, task.Status, task.StartDate, task.EndDate).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert task: %w", err)
	}
	return nil
}

func (r *TaskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task

	query := `
	SELECT 
		id, title, description, status,
		start_date, end_date, created_at, updated_at
	FROM tasks
	WHERE id = $1
	`
	err := r.DB.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.StartDate,
		&task.EndDate,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	return &task, nil
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
	query := `
		SELECT 
			id, title, description, status,
			start_date, end_date, created_at, updated_at
		FROM tasks
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}
	defer rows.Close()

	tasks := []model.Task{}

	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.StartDate,
			&task.EndDate,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return tasks, nil
}

func (r *TaskRepository) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
