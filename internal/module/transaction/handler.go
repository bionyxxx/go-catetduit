package transaction

import (
	"catetduit/internal/helper"
	"catetduit/internal/middleware"
	"encoding/json"
	"fmt"
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

func (h *Handler) GetTransactionsByUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*helper.JWTClaims)
	if !ok {
		err := helper.ResponseUnauthorized(w, "Unauthorized access")
		if err != nil {
			panic(err.Error())
		}
		return
	}

	var req GetTransactionsByUserIDRequest
	req.UserID = claims.UserID
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	limit := uint(10) // default
	offset := uint(0) // default

	if limitStr != "" {
		if val, err := strconv.ParseUint(limitStr, 10, 32); err == nil {
			limit = uint(val)
		}
	}

	if offsetStr != "" {
		if val, err := strconv.ParseUint(offsetStr, 10, 32); err == nil {
			offset = uint(val)
		}
	}

	// startDate and endDate can be empty timestamp
	var startDate, endDate int64
	if startDateStr != "" {
		if val, err := strconv.ParseInt(startDateStr, 10, 64); err == nil {
			startDate = val
		}
	}

	if endDateStr != "" {
		if val, err := strconv.ParseInt(endDateStr, 10, 64); err == nil {
			endDate = val
		}
	}

	req.Limit = limit
	req.Offset = offset
	req.StartDate = startDate
	req.EndDate = endDate
	transactionsResp, err := h.service.GetTransactionsByUserID(&req)
	if err != nil {
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			panic(err.Error())
		}
		return
	}

	err = helper.ResponseOKWithData(w, "Retrieval successful", transactionsResp)
	if err != nil {
		panic(err.Error())
	}
}

func (h *Handler) GetTransactionSummaryByUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*helper.JWTClaims)
	if !ok {
		err := helper.ResponseUnauthorized(w, "Unauthorized access")
		if err != nil {
			panic(err.Error())
		}
		return
	}

	userId := claims.UserID

	summaryResp, err := h.service.GetTransactionSummaryByUserID(userId)
	if err != nil {
		err := helper.ResponseInternalServerError(w, "An error occurred, please try again.", err.Error())
		if err != nil {
			panic(err.Error())
		}
		return
	}

	err = helper.ResponseOKWithData(w, "Retrieval successful", summaryResp)
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
