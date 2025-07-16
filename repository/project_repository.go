package repository

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq" // PostgreSQL driver

	"github.com/ChileKasoka/construction-app/model"
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) GetAll() ([]model.Project, error) {
	query := `
	SELECT id, name, description, start_date, end_date, status, created_at, updated_at
	FROM projects
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []model.Project
	for rows.Next() {
		var project model.Project
		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
			&project.StartDate,
			&project.EndDate,
			&project.Status,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) GetByID(id int) (*model.Project, error) {
	query := `
SELECT id, name, description, start_date, end_date, status
FROM projects
WHERE id = $1
`
	row := r.DB.QueryRow(query, id)
	var project model.Project
	err := row.Scan(&project.ID, &project.Name, &project.Description, &project.StartDate, &project.EndDate, &project.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("project not found")
		}
		return nil, err
	}

	return &project, nil
}

func (r *ProjectRepository) GetAllCount() (int, error) {
	query := `SELECT COUNT(*) FROM projects`
	var count int
	err := r.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ProjectRepository) Create(p *model.Project) (*model.Project, error) {
	query := `
	INSERT INTO projects (name, description, start_date, end_date, status)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`
	err := r.DB.QueryRow(query, p.Name, p.Description, p.StartDate, p.EndDate, p.Status).Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProjectRepository) Update(id int, p *model.Project) (*model.Project, error) {
	query := `
	UPDATE projects
	SET name = $1, description = $2, start_date = $3, end_date = $4, status = $5
	WHERE id = $6
	RETURNING id
	`
	err := r.DB.QueryRow(query, p.Name, p.Description, p.StartDate, p.EndDate, p.Status, id).Scan(&p.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("project not found")
		}
		return nil, err
	}
	return p, nil
}

func (r *ProjectRepository) Delete(id int) error {
	query := `
	DELETE FROM projects WHERE id = $1
	`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("project not found")
		}
		return err
	}
	return nil
}
