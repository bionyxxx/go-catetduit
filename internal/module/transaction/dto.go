package transaction

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

type CreateTransactionRequest struct {
	UserID          uint    `json:"user_id" validate:"required"`
	Type            string  `json:"type" validate:"required,oneof=credit debit"`
	Amount          float64 `json:"amount" validate:"required,gte=0"`
	Description     string  `json:"description" validate:"max=255"`
	TransactionDate int64   `json:"transaction_date" validate:"required,gte=0"`
}
