package category

import (
	"catetduit/internal/helper"
	"catetduit/internal/middleware"
	"net/http"
	"strconv"

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

}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {

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
