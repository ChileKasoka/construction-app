package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/ChileKasoka/construction-app/model"
)

type UserProjectRepository struct {
	DB *sql.DB
}

func NewUserProjectRepository(db *sql.DB) *UserProjectRepository {
	return &UserProjectRepository{DB: db}
}

func (r *UserProjectRepository) Create(userProject *model.UserProject) error {
	query := `
	INSERT INTO user_project (project_id, user_id)
	VALUES ($1, $2)
	RETURNING project_id, user_id
	`

	err := r.DB.QueryRow(query, userProject.ProjectID, userProject.UserID).
		Scan(&userProject.ProjectID, &userProject.UserID)

	if err != nil {
		return err
	}
	return nil
}

func (r *UserProjectRepository) CreateMany(projectID int, userIDs []int) error {
	if len(userIDs) == 0 {
		return nil
	}

	query := `INSERT INTO user_project (project_id, user_id) VALUES `
	values := []any{}
	placeholders := []string{}

	for i, userID := range userIDs {
		projectPos := i*2 + 1
		userPos := i*2 + 2
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", projectPos, userPos))
		values = append(values, projectID, userID)
	}

	fullQuery := query + strings.Join(placeholders, ", ")

	_, err := r.DB.Exec(fullQuery, values...)
	return err
}

func (r *UserProjectRepository) GetAll() ([]model.User, error) {
	query := `
	SELECT u.id, u.name, u.email, u.password, pr.name as project, u.role_id, r.name, r.description, u.created_at
	FROM user_project up
	JOIN users u ON up.user_id = u.id
    JOIN projects pr ON up.project_id = pr.id
	LEFT JOIN roles r ON u.role_id = r.id
	ORDER BY u.id
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User
		var role model.Role

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Project, &user.RoleID, &role.Name, &role.Description, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		role.ID = user.RoleID
		user.Role = &role

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserProjectRepository) GetByProjectID(projectID int) ([]model.ProjectUsers, error) {
	query := `
	SELECT u.id AS user_id, u.name AS user_name, u.email AS user_email, pr.id AS project_id, pr.name AS project, r.name AS role_name
	FROM user_project up
	JOIN users u ON up.user_id = u.id
	JOIN projects pr ON up.project_id = pr.id
	LEFT JOIN roles r ON u.role_id = r.id
	WHERE up.project_id = $1
	ORDER BY u.id
	`

	rows, err := r.DB.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.ProjectUsers

	for rows.Next() {
		var projectUser model.ProjectUsers

		err := rows.Scan(
			&projectUser.UserID,
			&projectUser.UserName,
			&projectUser.UserEmail,
			&projectUser.ProjectID,
			&projectUser.Project,
			&projectUser.RoleName,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, projectUser)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserProjectRepository) GetByUserID(userID int) ([]model.Project, error) {
	query := `
	SELECT p.id, p.name, p.description, p.start_date, p.end_date, p.created_at, p.updated_at
	FROM user_project up
	JOIN projects p ON up.project_id = p.id
	WHERE up.user_id = $1
	ORDER BY p.id
	`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project

	for rows.Next() {
		var project model.Project

		err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.StartDate, &project.EndDate, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
