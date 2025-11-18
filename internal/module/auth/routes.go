package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

// RegisterRoutes registers all auth-related routes
func RegisterRoutes(r chi.Router, validator *validator.Validate, authService *Service) {
	handler := NewHandler(authService, validator)

	r.Post("/auth/login", handler.Login)
	r.Post("/auth/register", handler.Register)
}
