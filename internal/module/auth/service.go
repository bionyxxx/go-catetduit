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

	_ = userData

	// TODO: Generate and return JWT token or session info

	return nil
}

func (s *Service) Register(name, phone, email, password string) error {
	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Create user model to repository
	newUser := &user.User{
		Name:     name,
		Phone:    phone,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil
}
