package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/ChileKasoka/construction-app/model"
)

type RolePermissionRepo struct {
	DB *sql.DB
}

func NewRolePerissionRepo(db *sql.DB) *RolePermissionRepo {
	return &RolePermissionRepo{DB: db}
}

func (r *RolePermissionRepo) HasPermission(roleID int, path string) (bool, error) {
	permissions, err := r.GetByRoleID(roleID)
	if err != nil {
		return false, err
	}

	for _, p := range permissions {
		if p.Path == path {
			return true, nil
		}
	}
	return false, nil
}

func matchPath(pattern, path string) bool {
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternParts) != len(pathParts) {
		return false
	}

	for i := 0; i < len(patternParts); i++ {
		if strings.HasPrefix(patternParts[i], ":") {
			continue // parameter placeholder matches anything
		}
		if patternParts[i] != pathParts[i] {
			return false
		}
	}
	return true
}

func (r *RolePermissionRepo) Create(roleID int, permissionIDs []int) error {
	query := `
		INSERT INTO role_permissions (role_id, permission_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`

	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, permissionID := range permissionIDs {
		_, err := stmt.Exec(roleID, permissionID)
		if err != nil {
			return fmt.Errorf("error assigning permission %d: %w", permissionID, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
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

type RolePermissionDisplay struct {
	ID             int    `json:"id"`
	RoleName       string `json:"role_name"`
	PermissionName string `json:"permission_name"`
	Path           string `json:"path"`
	Method         string `json:"method"`
}

func (r *RolePermissionRepo) GetAllRolePermissions() ([]RolePermissionDisplay, error) {
	query := `
		SELECT
			rp.id,
			r.name AS role_name,
			p.name AS permission_name,
			p.path,
			p.method
		FROM
			role_permissions rp
		JOIN roles r ON rp.role_id = r.id
		JOIN permissions p ON rp.permission_id = p.id
		ORDER BY r.name, p.name;
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query role-permissions: %w", err)
	}
	defer rows.Close()

	var result []RolePermissionDisplay

	for rows.Next() {
		var rp RolePermissionDisplay
		err := rows.Scan(&rp.ID, &rp.RoleName, &rp.PermissionName, &rp.Path, &rp.Method)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		result = append(result, rp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return result, nil
}

func (r *RolePermissionRepo) GetByUserID(id int) ([]*model.UserPermission, error) {
	query := `
		SELECT 
			u.id AS user_id,
			u.name AS user_name,
			r.name AS role_name,
			p.id AS permission_id,
			p.name AS permission_name,
			p.path,
			p.method
		FROM users u
		JOIN roles r ON u.role_id = r.id
		JOIN role_permissions rp ON rp.role_id = r.id
		JOIN permissions p ON rp.permission_id = p.id
		WHERE u.id = $1;
	`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var permissions []*model.UserPermission
	for rows.Next() {
		var up model.UserPermission
		err := rows.Scan(
			&up.UserID,
			&up.UserName,
			&up.RoleName,
			&up.PermissionID,
			&up.PermissionName,
			&up.Path,
			&up.Method,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		permissions = append(permissions, &up)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
