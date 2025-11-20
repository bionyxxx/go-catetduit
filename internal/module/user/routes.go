package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

// RegisterRoutes registers all user-related routes
func RegisterRoutes(r chi.Router, validator *validator.Validate, userService *Service) {
	handler := NewHandler(userService, validator)

	r.Get("/me", handler.Me)
	r.Patch("/change-password", handler.ChangePassword)
}
