package auth

import (
	"catetduit/internal/module/user"
	"errors"
)

var (
	// ErrInvalidCredentials is returned when the provided credentials are invalid
	ErrInvalidCredentials = errors.New("invalid credentials")

	// ErrUserAlreadyExists is returned when trying to register a user that already exists
	ErrUserAlreadyExists = errors.New("user already exists")

	// ErrUserNotFound is returned when a user is not found
	ErrUserNotFound = errors.New("user not found")
)

// Service represents the authentication service
type Service struct {
	userRepo user.Repository
}

// NewService creates a new authentication service
func NewService(userRepo user.Repository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

// Authenticate authenticates a user with the given email and password
func (s *Service) Authenticate(email, password string) error {
	userData, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return ErrUserNotFound
	}

	_ = userData // In a real implementation, you would verify the password here

	return nil
}

// Register registers a new user with the given details
func (s *Service) Register(name, email, password string) error {
	// Registration logic goes here
	return nil
}
