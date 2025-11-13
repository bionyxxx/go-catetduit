package auth

import "errors"

var (
	// ErrInvalidCredentials is returned when the provided credentials are invalid
	ErrInvalidCredentials = errors.New("invalid credentials")

	// ErrUserAlreadyExists is returned when trying to register a user that already exists
	ErrUserAlreadyExists = errors.New("user already exists")

	// ErrUserNotFound is returned when a user is not found
	ErrUserNotFound = errors.New("user not found")
)

// Service represents the authentication service
type Service struct{}

// NewService creates a new authentication service
func NewService() *Service {
	return &Service{}
}

// Authenticate authenticates a user with the given email and password
func (s *Service) Authenticate(email, password string) error {
	// Authentication logic goes here
	return nil
}

// Register registers a new user with the given details
func (s *Service) Register(name, email, password string) error {
	// Registration logic goes here
	return nil
}
