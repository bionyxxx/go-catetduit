package auth

import (
	"catetduit/internal/helper"
	"catetduit/internal/module/user"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
)

type Service struct {
	userRepo  user.Repository
	jwtHelper *helper.JWTHelper
}

func NewService(userRepo user.Repository, jwtHelper *helper.JWTHelper) *Service {
	return &Service{
		userRepo:  userRepo,
		jwtHelper: jwtHelper,
	}
}

func (s *Service) Authenticate(email, password string) (loginResp *LoginResponse, err error) {
	userData, err := s.userRepo.GetUserByEmail(email)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password)) != nil {
		return loginResp, ErrInvalidCredentials
	}

	accessToken, exp, err := s.jwtHelper.GenerateAccessToken(uint(userData.ID), userData.Email, userData.Name)
	if err != nil {
		return loginResp, err
	}

	refreshToken, err := s.jwtHelper.GenerateRefreshToken(uint(userData.ID), userData.Email, userData.Name)
	if err != nil {
		return loginResp, err
	}

	loginResp = &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    exp,
	}

	return loginResp, nil
}

func (s *Service) Register(name, phone, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
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
