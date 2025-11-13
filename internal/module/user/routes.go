package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// RegisterRoutes registers all user-related routes
func RegisterRoutes(r chi.Router, db *sqlx.DB) {
	handler := NewHandler(db)

	r.Get("/user", handler.GetUser)
}
