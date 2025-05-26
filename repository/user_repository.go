package repository

import (
	"database/sql"
	"errors"

	"github.com/ChileKasoka/construction-app/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *model.RegisterRequest) error {
	query := `
	INSERT INTO users (name, email, password, role_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
	`

	err := r.DB.QueryRow(query, user.Name, user.Email, user.Password, user.RoleID).
		Scan(&user.ID, &user.CreatedAt)

	return err
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	query := `
	SELECT u.id, u.name, u.email, u.password, u.role_id, r.name, r.description, u.created_at
	FROM users u
	LEFT JOIN roles r ON u.role_id = r.id
	ORDER BY u.id
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User
		var role model.Role

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &role.Name, &role.Description, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		role.ID = user.RoleID
		user.Role = &role

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetAllCount() (int, error) {
	query := `SELECT COUNT(*) FROM users`
	var count int
	err := r.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *UserRepository) GetByID(id int) (*model.User, error) {
	query := `
	SELECT u.id, u.name, u.email, u.password, u.role_id, r.name, r.description, u.created_at
	FROM users u
	LEFT JOIN roles r ON u.role_id = r.id
	WHERE u.id = $1
	`
	row := r.DB.QueryRow(query, id)

	var user model.User
	var role model.Role

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &role.Name, &role.Description, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	role.ID = user.RoleID
	user.Role = &role

	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := `
	SELECT u.id, u.name, u.email, u.password, u.role_id, r.name, r.description, u.created_at
	FROM users u
	LEFT JOIN roles r ON u.role_id = r.id
	WHERE u.email = $1
	`
	row := r.DB.QueryRow(query, email)

	var user model.User
	var role model.Role

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &role.Name, &role.Description, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	role.ID = user.RoleID
	user.Role = &role

	return &user, nil
}

func (r *UserRepository) Update(user model.User) error {
	query := `
	UPDATE users SET name=$1, email=$2, password=$3, role_id=$4 WHERE id=$5
	`
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.RoleID, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := r.DB.Exec(query, id)
	return err
}
