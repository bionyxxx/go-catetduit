package category

type CategoryResponse struct {
	ID        *uint  `json:"id,omitempty"`
	UserID    *uint  `json:"user_id,omitempty"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CategoryLoadMoreResponse struct {
	Categories []*CategoryResponse `json:"categories"`
	HasMore    bool                `json:"has_more"`
}

type GetCategoriesByUserIDRequest struct {
	UserID uint  `json:"user_id" validate:"required"`
	Limit  *uint `json:"limit" validate:"gte=0"`
	Offset *uint `json:"offset" validate:"gte=0"`
}

type CreateCategoryRequest struct {
	UserID uint   `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required,max=100"`
}
