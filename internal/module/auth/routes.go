package auth

import (
	"catetduit/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func RegisterRoutes(r chi.Router, validator *validator.Validate, authService *Service, oauth2Config *config.OAuth2Config) {
	handler := NewHandler(authService, validator)

	r.Post("/auth/login", handler.Login)
	r.Post("/auth/refresh", handler.Refresh)
	r.Post("/auth/register", handler.Register)
	r.Post("/auth/logout", handler.Logout)
}
