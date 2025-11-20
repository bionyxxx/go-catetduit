package user

import (
	"catetduit/internal/helper"
	"catetduit/internal/middleware"
	"encoding/json"
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

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*helper.JWTClaims)
	if !ok {
		err := helper.ResponseUnauthorized(w, "Unauthorized access")
		if err != nil {
			panic(err.Error())
		}
		return
	}

	userId := claims.UserID

	user, err := h.service.GetUserByID(userId)

	err = helper.ResponseOKWithData(w, "Retrieval successful", user)

	if err != nil {
		panic(err.Error())
	}
}

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*helper.JWTClaims)

	if !ok {
		err := helper.ResponseUnauthorized(w, "Unauthorized access")
		if err != nil {
			panic(err.Error())
		}
		return
	}

	userId := claims.UserID

	// Register custom validation for old password
	err := h.validator.RegisterValidation("old_password", func(fl validator.FieldLevel) bool {
		oldPassword := fl.Field().String()
		valid, err := h.oldPasswordValidator(userId, oldPassword)
		if err != nil || !valid {
			return false
		}
		return true
	})
	if err != nil {
		panic("Error registering custom validation: " + err.Error())
	}

	var req ChangePasswordRequest
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

	err = h.service.ChangePassword(userId, req.NewPassword)

	if err != nil {
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	err = helper.ResponseOK(w, "Password changed successfully")

	if err != nil {
		fmt.Println("Error sending response:", err)
	}

}

// oldPasswordValidator checks if the provided old password matches the stored password for a user.
func (h *Handler) oldPasswordValidator(userID uint, oldPassword string) (bool, error) {
	return h.service.CheckOldPassword(userID, oldPassword)
}
