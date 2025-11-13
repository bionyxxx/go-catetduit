package auth

import "net/http"

type Handler struct{}

// NewHandler creates a new user handler
func NewHandler() *Handler {
	return &Handler{}
}

// Login handles user login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

}

// Register handles user registration
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

}
