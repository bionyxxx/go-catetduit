package user

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

var tableName = "users"

type repositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) CreateUser(user *User) (*User, error) {
	query := `
		INSERT INTO users (name, phone, email, password) 
		VALUES (:name, :phone, :email, :password)
		RETURNING id, name, phone, email, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	if !rows.Next() {
		return nil, errors.New("user not found")
	}

	var createdUser User
	if err = rows.StructScan(&createdUser); err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (r *repositoryImpl) GetUserByID(id uint) (*User, error) {
	var user User
	err := r.db.Get(&user, "SELECT id, name, phone, email, password FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repositoryImpl) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.db.Get(&user, "SELECT id, name, phone, email, password FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
