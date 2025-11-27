package transaction

import "time"

type TransactionResponse struct {
	ID              *uint   `json:"id,omitempty"`
	UserID          *uint   `json:"user_id,omitempty"`
	Amount          float64 `json:"amount"`
	Type            string  `json:"type"` // e.g., "credit" or "debit"
	Description     string  `json:"description"`
	TransactionDate int64   `json:"transaction_date"`
	CreatedAt       int64   `json:"created_at"`
	UpdatedAt       int64   `json:"updated_at"`
}

type TransactionSummaryResponse struct {
	TotalCredit float64 `json:"total_credit"`
	TotalDebit  float64 `json:"total_debit"`
	Balance     float64 `json:"balance"`
}

type TransactionLoadMoreResponse struct {
	Transactions []*TransactionResponse `json:"transactions"`
	HasMore      bool                   `json:"has_more"`
}

type GetTransactionsByUserIDRequest struct {
	UserID    uint       `json:"user_id" validate:"required"`
	Limit     uint       `json:"limit" validate:"gte=0"`
	Offset    uint       `json:"offset" validate:"gte=0"`
	StartDate *time.Time `json:"start_date" validate:"gte=0"`
	EndDate   *time.Time `json:"end_date" validate:"gte=0"`
}

type CreateTransactionRequest struct {
	UserID          uint    `json:"user_id" validate:"required"`
	Type            string  `json:"type" validate:"required,oneof=credit debit"`
	Amount          float64 `json:"amount" validate:"required,gte=0"`
	Description     string  `json:"description" validate:"max=255"`
	TransactionDate int64   `json:"transaction_date" validate:"required,gte=0"`
}
