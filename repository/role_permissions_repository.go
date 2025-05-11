package repository

import (
	"database/sql"

	"github.com/ChileKasoka/construction-app/model"
)

type RolePermissionRepo struct {
	DB *sql.DB
}

func (r *RolePermissionRepo) Create(roleID, permissionID int) error {
	_, err := r.DB.Exec(`
		INSERT INTO role_permissions (role_id, permission_id)
		VALUES ($1, $2)
		ON CONFLICT (role_id, permission_id) DO NOTHING
	`, roleID, permissionID)
	return err
}

func (r *RolePermissionRepo) Delete(roleID, permissionID int) error {
	_, err := r.DB.Exec(`
		DELETE FROM role_permissions
		WHERE role_id = $1 AND permission_id = $2
	`, roleID, permissionID)
	return err
}

func (r *RolePermissionRepo) GetByRoleID(roleID int) ([]model.Permission, error) {
	rows, err := r.DB.Query(`
		SELECT p.id, p.name, p.path, p.method
		FROM role_permissions rp
		JOIN permissions p ON rp.permission_id = p.id
		WHERE rp.role_id = $1
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
	return permissions, nil
}
