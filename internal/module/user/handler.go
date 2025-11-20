package user

import (
	"catetduit/internal/helper"
	"catetduit/internal/middleware"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	claims, err := r.Context().Value(middleware.UserClaimsKey).(*helper.JWTClaims)
	if !err {
		err := helper.ResponseUnauthorized(w, "Unauthorized access")
		if err != nil {
			panic(err.Error())
		}
		return
	}

	userId := claims.UserID

	user, err := h.service.GetUserByID(userId)

	//TODO : fix me response

	err = helper.ResponseOKWithData(w, "Retrieval successful", user)

	if err != nil {
		panic(err.Error())
	}
}
