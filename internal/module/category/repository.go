package category

type Repository interface {
	GetCategoryByID(userID, id uint) (*Category, error)
	CreateCategory(category *Category) (*Category, error)
	GetCategoriesByUserID(userID uint, limit, offset *uint) ([]*Category, error)
}
