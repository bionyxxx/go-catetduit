package category

import (
	"github.com/jmoiron/sqlx"
)

type repositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) GetCategoryByID(id uint) (*Category, error) {
	// Implementation goes here
	return nil, nil
}

func (r *repositoryImpl) CreateCategory(category *Category) (*Category, error) {
	// Implementation goes here
	return nil, nil
}

func (r *repositoryImpl) GetCategoriesByUserID(userID uint, limit, offset *uint) ([]*Category, error) {
	query := "SELECT id, user_id, name, created_at, updated_at FROM categories WHERE user_id = $1"
	args := []interface{}{userID}

	var categories []*Category
	err := r.db.Select(&categories, query, args...)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
