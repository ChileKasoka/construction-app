package repository

import (
	"database/sql"

	"github.com/ChileKasoka/construction-app/model"
)

type UserTaskRepository struct {
	DB *sql.DB
}

func NewUserTaskRepository(db *sql.DB) *UserTaskRepository {
	return &UserTaskRepository{DB: db}
}

func (r *UserTaskRepository) Create(userTask *model.UserTask) error {
	query := `
	INSERT INTO user_tasks (task_id, user_id)
	VALUES ($1, $2)
	RETURNING task_id, user_id
	`

	err := r.DB.QueryRow(query, userTask.TaskID, userTask.UserID).
		Scan(&userTask.TaskID, &userTask.UserID)

	if err != nil {
		return err
	}
	return nil
}

func (r *UserTaskRepository) GetAll() ([]model.UserTask, error) {
	query := `
	SELECT task_id, user_id
	FROM user_tasks
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userTasks []model.UserTask
	for rows.Next() {
		var userTask model.UserTask
		err := rows.Scan(&userTask.TaskID, &userTask.UserID)
		if err != nil {
			return nil, err
		}
		userTasks = append(userTasks, userTask)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userTasks, nil
}

func (r *UserTaskRepository) GetByUserID(userID int) ([]model.UserTask, error) {
	query := `
	SELECT task_id, user_id
	FROM user_tasks
	WHERE user_id = $1
	`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userTasks []model.UserTask
	for rows.Next() {
		var userTask model.UserTask
		err := rows.Scan(&userTask.TaskID, &userTask.UserID)
		if err != nil {
			return nil, err
		}
		userTasks = append(userTasks, userTask)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userTasks, nil
}

func (r *UserTaskRepository) UnassignUser(taskID, userID int) error {
	query := `DELETE FROM task_assignees WHERE task_id = $1 AND user_id = $2`
	_, err := r.DB.Exec(query, taskID, userID)
	return err
}
