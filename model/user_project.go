package model

type UserProject struct {
	ProjectID int `json:"project_id"`
	UserID    int `json:"user_id"`
}

type ProjectUsers struct {
	ProjectID int     `json:"project_id"`
	UserID    int     `json:"user_id"`
	UserName  string  `json:"user_name"`
	UserEmail string  `json:"user_email"`
	Project   string  `json:"project"`
	RoleName  *string `json:"role_name"`
}
