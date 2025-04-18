package repository

import (
	"errors"

	"github.com/ChileKasoka/construction-app/model"
)

type UserRepository struct{}

var users = []model.User{
	{ID: 1, Name: "Admin", Email: "admin@site.com", Role: "admin"},
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	return users, nil
}

func (r *UserRepository) GetByID(id int) (*model.User, error) {
	for _, u := range users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Update(id int, data map[string]interface{}) (*model.User, error) {
	for i, u := range users {
		if u.ID == id {
			if name, ok := data["name"].(string); ok {
				users[i].Name = name
			}
			if email, ok := data["email"].(string); ok {
				users[i].Email = email
			}
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *UserRepository) Delete(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
