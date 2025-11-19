package auth

import (
	"catetduit/internal/helper"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service   *Service
	validator *validator.Validate
}

func NewHandler(service *Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := helper.ResponseBadRequest(w, "Invalid request payload", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	if err := h.validator.Struct(req); err != nil {
		errDetails := helper.FormatValidationErrors(err)
		err := helper.ResponseUnprocessableEntity(w, "Validation failed", errDetails)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	loginResp, err := h.service.Authenticate(req.Email, req.Password)

	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			err := helper.ResponseUnauthorized(w, "Invalid email or password")
			if err != nil {
				fmt.Println("Error sending response:", err)
			}
			return
		}
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	err = helper.ResponseOKWithData(w, "Login successful", loginResp)
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := helper.ResponseBadRequest(w, "Invalid request payload", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	if err := h.validator.Struct(req); err != nil {
		errDetails := helper.FormatValidationErrors(err)
		err := helper.ResponseUnprocessableEntity(w, "Validation failed", errDetails)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	err := h.service.Register(req.Name, req.Phone, req.Email, req.Password)

	if err != nil {
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	err = helper.ResponseCreated(w, "Registration successful", nil)
	if err != nil {
		fmt.Println("Error sending response:", err)
	}

}
