package user

import (
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers all user-related routes
func RegisterRoutes(r chi.Router, userService *Service) {
	handler := NewHandler(userService)

	r.Get("/me", handler.Me)
}
