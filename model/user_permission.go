package model

type UserPermission struct {
	UserID         int
	UserName       string
	RoleName       string
	PermissionID   int
	PermissionName string
	Path           string
	Method         string
}
