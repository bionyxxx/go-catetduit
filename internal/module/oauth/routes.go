package oauth

import (
	"catetduit/internal/config"
	"catetduit/internal/module/auth"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func RegisterRoutes(r chi.Router, validator *validator.Validate, authService *auth.Service, oauth2Service *Service, oauth2Config *config.OAuth2Config) {
	handler := NewHandler(oauth2Service, validator)

	r.Get("/google", handler.Google)
	r.Get("/google/callback", handler.GoogleCallback)
}
