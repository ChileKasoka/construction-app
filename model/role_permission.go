package model

type RolePermission struct {
	ID           int `json:"id"`
	RoleID       int `json:"role_id"`
	PermissionID int `json:"permission_id"`
}
