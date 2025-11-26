package oauth

import (
	"catetduit/internal/config"
	"catetduit/internal/helper"
	"catetduit/internal/module/auth"
	"catetduit/internal/module/user"
)

type Service struct {
	oauth2Config *config.OAuth2Config
	userRepo     user.Repository
	jwtHelper    *helper.JWTHelper
}

func NewService(auth2Config *config.OAuth2Config, userRepo user.Repository, helper *helper.JWTHelper) *Service {
	return &Service{
		oauth2Config: auth2Config,
		userRepo:     userRepo,
		jwtHelper:    helper,
	}
}

func (s *Service) Google(info *GoogleUserInfo) (*auth.LoginResponse, error) {
	userData, err := s.userRepo.GetUserByEmail(info.Email)

	if err != nil {
		// User not found, create new user
		newUser := &user.User{
			Name:  info.Name,
			Email: info.Email,
		}

		userData, err = s.userRepo.CreateUser(newUser)
		if err != nil {
			return nil, err
		}
	}

	accessToken, exp, err := s.jwtHelper.GenerateAccessToken(uint(userData.ID), userData.Email, userData.Name)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtHelper.GenerateRefreshToken(uint(userData.ID), userData.Email, userData.Name)
	if err != nil {
		return nil, err
	}

	loginResp := &auth.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    exp,
	}

	return loginResp, nil
}
