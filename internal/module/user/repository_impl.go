package user

import "github.com/jmoiron/sqlx"

type repositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) GetUserByID(id int) (*User, error) {
	var user User
	err := r.db.Get(&user, "SELECT id, name, email, age FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repositoryImpl) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.db.Get(&user, "SELECT id, name, email, age FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
