package user

type Service struct {
	repo Repository
}

// NewService creates a new user service
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetUser retrieves user information
func (s *Service) GetUser(id int) (*User, error) {
	return s.repo.GetUserByID(id)
}
