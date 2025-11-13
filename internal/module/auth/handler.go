package auth

import (
	"catetduit/internal/helper"
	"net/http"
)

type Handler struct {
	service *Service
}

// NewHandler creates a new user handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Login handles user login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	_ = h.service.Authenticate(email, password)

	err := helper.ResponseOK(w, "Login successful")
	if err != nil {
		return
	}
}

// Register handles user registration
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

}
