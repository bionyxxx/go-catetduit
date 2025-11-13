package auth

import (
	"catetduit/internal/module/user"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
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
		return ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password)); err != nil {
		return ErrInvalidCredentials
	}

	_ = userData // In a real implementation, you would verify the password here

	return nil
}

// Register registers a new user with the given details
func (s *Service) Register(name, email, password string, age int) error {

	return nil
}
