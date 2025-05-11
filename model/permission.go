package model

type Permission struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
