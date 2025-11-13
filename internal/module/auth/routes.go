package auth

import "github.com/go-chi/chi/v5"

// RegisterRoutes registers all auth-related routes
func RegisterRoutes(r chi.Router) {
	handler := NewHandler()

	r.Post("/auth/login", handler.Login)
	r.Post("/auth/register", handler.Register)
}
