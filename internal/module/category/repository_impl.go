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

func (r *repositoryImpl) GetCategoryByID(userID, id uint) (*Category, error) {
	query := "SELECT id, user_id, name, created_at, updated_at FROM categories WHERE user_id = $1 AND id = $2"

	var category Category
	err := r.db.Get(&category, query, userID, id)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *repositoryImpl) CreateCategory(category *Category) (*Category, error) {
	query := `
			INSERT INTO categories (user_id, name)
			VALUES (:user_id, :name)
			RETURNING id, user_id, name, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, category)
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
		return nil, err
	}

	var createdCategory Category
	err = rows.StructScan(&createdCategory)
	if err != nil {
		return nil, err
	}

	return &createdCategory, nil
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
