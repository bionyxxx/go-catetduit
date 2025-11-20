package user

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUserByID(id uint) (*UserResponse, error) {
	userData, err := s.repo.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	userResp := &UserResponse{
		ID:        &userData.ID,
		Name:      userData.Name,
		Phone:     userData.Phone,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt.Unix(),
		UpdatedAt: userData.UpdatedAt.Unix(),
	}

	return userResp, nil
}
