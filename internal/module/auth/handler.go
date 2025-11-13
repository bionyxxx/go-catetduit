package auth

import (
	"catetduit/internal/helper"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service   *Service
	validator *validator.Validate
}

// NewHandler creates a new user handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service:   service,
		validator: validator.New(),
	}
}

// Login handles user login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	//validate input using go-playground/validator
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := helper.ResponseBadRequest(w, "Invalid request payload", err.Error())
		if err != nil {
			return
		}
		return
	}

	if err := h.validator.Struct(req); err != nil {
		details := helper.FormatValidationErrors(err)
		err := helper.ResponseUnprocessableEntity(w, "Validation failed", details)
		if err != nil {
			return
		}
		return
	}

	email := req.Email
	password := req.Password

	err := h.service.Authenticate(email, password)

	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			err := helper.ResponseUnauthorized(w, "Invalid email or password")
			if err != nil {
				return
			}
			return
		}
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			return
		}
		return
	}

	err = helper.ResponseOK(w, "Login successful")
	if err != nil {
		return
	}
}

// Register handles user registration
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

}
