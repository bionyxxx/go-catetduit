package auth

import (
	"catetduit/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

// RegisterRoutes registers all auth-related routes
func RegisterRoutes(r chi.Router, validator *validator.Validate, authService *Service, oauth2Config *config.OAuth2Config) {
	handler := NewHandler(authService, validator)

	r.Post("/auth/login", handler.Login)
	r.Post("/auth/refresh", handler.Refresh)
	r.Post("/auth/register", handler.Register)
	r.Get("/auth/google/login", handler.GoogleLogin(oauth2Config))
	//r.Get("/auth/google/callback", handler.GoogleCallback(oauth2Config))
}
