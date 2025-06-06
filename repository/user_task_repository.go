package repository

import (
	"database/sql"
	"fmt"
	"strings"

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

func (r *UserTaskRepository) CreateMany(taskID int, userIDs []int) error {
	if len(userIDs) == 0 {
		return nil
	}

	query := `INSERT INTO user_task (task_id, user_id) VALUES `
	values := []any{}
	placeholders := []string{}

	for i, userID := range userIDs {
		taskPos := i*2 + 1
		userPos := i*2 + 2
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", taskPos, userPos))
		values = append(values, taskID, userID) // Repeats taskID for each user
	}

	fullQuery := query + strings.Join(placeholders, ", ")

	_, err := r.DB.Exec(fullQuery, values...)
	return err
}

func (r *UserTaskRepository) GetAll() ([]model.UserTask, error) {
	query := `
	SELECT task_id, user_id
	FROM user_task
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
		SELECT t.id AS task_id, t.title AS task_title, t.description AS task_description, ut.user_id
		FROM user_task ut
		JOIN tasks t ON ut.task_id = t.id
		WHERE ut.user_id = $1
		ORDER BY t.id
	`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userTasks []model.UserTask
	for rows.Next() {
		var userTask model.UserTask
		err := rows.Scan(&userTask.TaskID, &userTask.Title, &userTask.Description, &userTask.UserID)
		if err != nil {
			return nil, err
		}
		userTasks = append(userTasks, userTask)
		fmt.Printf("UserTask: %+v\n", userTask) // Debugging output
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userTasks, nil
}

func (r *UserTaskRepository) UnassignUser(taskID, userID int) error {
	query := `DELETE FROM user_task WHERE task_id = $1 AND user_id = $2`
	_, err := r.DB.Exec(query, taskID, userID)
	return err
}
