package model

type UserTask struct {
	TaskID      int    `json:"task_id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
