package repository

import (
	"database/sql"

	"github.com/ChileKasoka/construction-app/model"
)

type RoleRepository struct {
	DB *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) Create(role *model.Role) error {
	query := `
	INSERT INTO roles(name, description)
	VALUES ($1, $2)
	RETURNING id, created_at
	`

	err := r.DB.QueryRow(query, role.Name, role.Description).
		Scan(&role.ID, &role.CreatedAt)
	return err
}

func (r *RoleRepository) FindByName(name string) (*model.Role, error) {
	query := `
	SELECT id, name
	FROM roles
	WHERE name = $1
	`
	row := r.DB.QueryRow(query, name)
	var role model.Role
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) GetAll() ([]model.Role, error) {
	query := `
	SELECT id, name, description, created_at
	FROM roles
	ORDER BY id
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []model.Role

	for rows.Next() {
		var role model.Role
		err := rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepository) GetByID(id int) (*model.Role, error) {
	query := `
	SELECT id, name, description, created_at
	FROM roles
	WHERE id = $1
	`
	row := r.DB.QueryRow(query, id)
	var role model.Role
	err := row.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) Update(role *model.Role) error {
	query := `
	UPDATE roles
	SET name = $1, description = $2
	WHERE id = $3
	`
	_, err := r.DB.Exec(query, role.Name, role.Description, role.ID)
	return err
}

func (r *RoleRepository) Delete(id int) error {
	query := `
	DELETE FROM roles
	WHERE id = $1
	`
	_, err := r.DB.Exec(query, id)
	return err
}
func (r *RoleRepository) GetByUserID(userID int) (*model.Role, error) {
	query := `
	SELECT r.id, r.name, r.description, r.created_at
	FROM roles r
	JOIN users u ON u.role_id = r.id
	WHERE u.id = $1
	`
	row := r.DB.QueryRow(query, userID)
	var role model.Role
	err := row.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &role, nil
}
