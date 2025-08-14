package repository

import (
	"database/sql"

	"github.com/ChileKasoka/construction-app/model"
)

type PermissionRepository struct {
	DB *sql.DB
}

func NewPermissionRepository(db *sql.DB) *PermissionRepository {
	return &PermissionRepository{DB: db}
}

func (r *PermissionRepository) Create(permission *model.Permission) error {
	query := `INSERT INTO permissions (name, path, method) VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRow(query, permission.Name, permission.Path, permission.Method).Scan(&permission.ID)
}

func (r *PermissionRepository) GetAll() ([]model.Permission, error) {
	rows, err := r.DB.Query(`SELECT id, name, path, method FROM permissions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []model.Permission
	for rows.Next() {
		var p model.Permission
		err := rows.Scan(&p.ID, &p.Name, &p.Path, &p.Method)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, p)
	}
	return permissions, nil
}

func (r *PermissionRepository) GetByID(id int64) (*model.Permission, error) {
	var p model.Permission
	err := r.DB.QueryRow(`SELECT id, name, path, method FROM permissions WHERE id = $1`, id).
		Scan(&p.ID, &p.Name, &p.Path, &p.Method)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PermissionRepository) GetUnassignedByRoleID(roleID int) ([]model.Permission, error) {
	rows, err := r.DB.Query(`
		SELECT p.id, p.name, p.path, p.method
		FROM permissions p
		WHERE p.id NOT IN (
			SELECT permission_id
			FROM role_permissions
			WHERE role_id = $1
		)
		ORDER BY p.id
	`, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []model.Permission
	for rows.Next() {
		var p model.Permission
		if err := rows.Scan(&p.ID, &p.Name, &p.Path, &p.Method); err != nil {
			return nil, err
		}
		permissions = append(permissions, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *PermissionRepository) Update(p *model.Permission) error {
	_, err := r.DB.Exec(`UPDATE permissions SET name = $1, path = $2, method = $3 WHERE id = $4`,
		p.Name, p.Path, p.Method, p.ID)
	return err
}

func (r *PermissionRepository) Delete(id int64) error {
	_, err := r.DB.Exec(`DELETE FROM permissions WHERE id = $1`, id)
	return err
}
