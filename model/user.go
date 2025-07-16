package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	RoleID    int       `json:"role_id"`
	Project   string    `json:"project,omitempty"` // Project name, if applicable
	Role      *Role     `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterRequest struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	// CreatedBy string    `json:"created_by,omitempty"` // User who created this record
}
