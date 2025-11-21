package transaction

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

func (h *Handler) GetTransactionsByUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*helper.JWTClaims)
	if !ok {
		err := helper.ResponseUnauthorized(w, "Unauthorized access")
		if err != nil {
			panic(err.Error())
		}
		return
	}

	userId := claims.UserID

	transactions, err := h.service.GetTransactionsByUserID(userId)
	if err != nil {
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			panic(err.Error())
		}
		return
	}

	err = helper.ResponseOKWithData(w, "Retrieval successful", transactions)
	if err != nil {
		panic(err.Error())
	}
}

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*helper.JWTClaims)
	if !ok {
		err := helper.ResponseUnauthorized(w, "Unauthorized access")
		if err != nil {
			panic(err.Error())
		}
		return
	}

	var req CreateTransactionRequest
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

	transactionResp, err := h.service.CreateTransaction(&req)

	if err != nil {
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	err = helper.ResponseCreated(w, "Transaction created successfully", transactionResp)
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}
