package model

import "time"

type Permission struct {
	ID        int       `json:"id"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
