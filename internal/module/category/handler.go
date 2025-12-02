package category

import (
	"catetduit/internal/helper"
	"catetduit/internal/middleware"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (h *Handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*helper.JWTClaims)
	if !ok {
		err := helper.ResponseUnauthorized(w, "Unauthorized access")
		if err != nil {
			panic(err.Error())
		}
		return
	}

	categoryIDStr := chi.URLParam(r, "id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		err := helper.ResponseBadRequest(w, "Invalid category ID", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	req := GetCategoryRequest{
		UserID: claims.UserID,
		ID:     uint(categoryID),
	}

	category, err := h.service.GetCategoryUserByID(&req)
	if err != nil {
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			panic(err.Error())
		}
		return
	}

	err = helper.ResponseOKWithData(w, "Category retrieved successfully", category)
	if err != nil {
		panic(err.Error())
	}
}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*helper.JWTClaims)
	if !ok {
		err := helper.ResponseUnauthorized(w, "Unauthorized access")
		if err != nil {
			panic(err.Error())
		}
		return
	}

	var req CreateCategoryRequest
	req.UserID = claims.UserID
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

	category, err := h.service.CreateCategory(&req)
	if err != nil {
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			panic(err.Error())
		}
		return
	}

	err = helper.ResponseCreated(w, "Category created successfully", category)
	if err != nil {
		panic(err.Error())
	}
}

func (h *Handler) GetCategoriesByUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*helper.JWTClaims)
	if !ok {
		err := helper.ResponseUnauthorized(w, "Unauthorized access")
		if err != nil {
			panic(err.Error())
		}
		return
	}

	var req GetCategoriesByUserIDRequest
	req.UserID = claims.UserID
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	if limitStr != "" {
		if val, err := strconv.ParseUint(limitStr, 10, 32); err == nil {
			limit := uint(val)
			req.Limit = &limit
		}
	}

	if offsetStr != "" {
		if val, err := strconv.ParseUint(offsetStr, 10, 32); err == nil {
			offset := uint(val)
			req.Offset = &offset
		}
	}

	categories, err := h.service.GetCategoriesByUserID(&req)
	if err != nil {
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			panic(err.Error())
		}
		return
	}

	err = helper.ResponseOKWithData(w, "Categories retrieved successfully", categories)
	if err != nil {
		panic(err.Error())
	}
}
