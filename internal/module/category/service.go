package category

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetCategoryByID(id uint) (*CategoryResponse, error) {
	return nil, nil
}

func (s *Service) CreateCategory(categoryRequest *CreateCategoryRequest) (*CategoryResponse, error) {
	//TODO: implement create category
	return nil, nil
}

func (s *Service) GetCategoriesByUserID(req *GetCategoriesByUserIDRequest) ([]*CategoryLoadMoreResponse, error) {
	var limit, offset *uint
	var hasMore bool = false

	if req.Limit != nil {
		limit = req.Limit
	}
	if req.Offset != nil {
		offset = req.Offset
	}

	categories, err := s.repo.GetCategoriesByUserID(req.UserID, limit, offset)
	if err != nil {
		return nil, err
	}

	if limit != nil && uint(len(categories)) > *limit {
		hasMore = true
		categories = categories[:*limit]
	}

	var categoryResponses []*CategoryResponse
	for _, category := range categories {
		categoryResponse := &CategoryResponse{
			ID:        &category.ID,
			UserID:    &category.UserID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt.Unix(),
			UpdatedAt: category.UpdatedAt.Unix(),
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}

	return []*CategoryLoadMoreResponse{
		{
			Categories: categoryResponses,
			HasMore:    hasMore,
		},
	}, nil
}
