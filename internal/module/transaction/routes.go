package transaction

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func RegisterRoutes(r chi.Router, validator *validator.Validate, transactionService *Service) {
	handler := NewHandler(transactionService, validator)

	r.Get("/transactions", handler.GetTransactionsByUser)
	r.Post("/transactions", handler.CreateTransaction)
}
