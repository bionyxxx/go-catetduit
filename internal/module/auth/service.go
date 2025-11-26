package auth

import (
	"catetduit/internal/config"
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
	userRepo     user.Repository
	jwtHelper    *helper.JWTHelper
	oauth2Config *config.OAuth2Config
}

func NewService(userRepo user.Repository, jwtHelper *helper.JWTHelper, auth2Config config.OAuth2Config) *Service {
	return &Service{
		userRepo:     userRepo,
		jwtHelper:    jwtHelper,
		oauth2Config: &auth2Config,
	}
}

func (s *Service) Authenticate(email, password string) (*LoginResponse, error) {
	userData, err := s.userRepo.GetUserByEmail(email)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password)) != nil {
		return nil, ErrInvalidCredentials
	}

	accessToken, exp, err := s.jwtHelper.GenerateAccessToken(uint(userData.ID), userData.Email, userData.Name)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtHelper.GenerateRefreshToken(uint(userData.ID), userData.Email, userData.Name)
	if err != nil {
		return nil, err
	}

	loginResp := &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    exp,
	}

	return loginResp, nil
}

func (s *Service) RefreshToken(refreshToken string) (*LoginResponse, error) {
	claims, err := s.jwtHelper.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	accessToken, exp, err := s.jwtHelper.GenerateAccessToken(claims.UserID, claims.Email, claims.Name)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.jwtHelper.GenerateRefreshToken(claims.UserID, claims.Email, claims.Name)
	if err != nil {
		return nil, err
	}

	loginResp := &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresAt:    exp,
	}

	return loginResp, nil
}

func (s *Service) Register(registerRequest *RegisterRequest) (*user.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &user.User{
		Name:     registerRequest.Name,
		Phone:    registerRequest.Phone,
		Email:    registerRequest.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := s.userRepo.CreateUser(newUser)

	if err != nil {
		return nil, err
	}

	userResp := &user.UserResponse{
		Name:      createdUser.Name,
		Phone:     createdUser.Phone,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt.Unix(),
		UpdatedAt: createdUser.UpdatedAt.Unix(),
	}

	return userResp, nil
}
