package transaction

import "time"

const TransactionTypeCredit = "credit"
const TransactionTypeDebit = "debit"

type Transaction struct {
	ID              uint      `json:"id" db:"id"`
	UserID          uint      `json:"user_id" db:"user_id"`
	Amount          float64   `json:"amount" db:"amount"`
	Type            string    `json:"type" db:"type"` // e.g., "credit" or "debit"
	Description     string    `json:"description" db:"description"`
	TransactionDate time.Time `json:"transaction_date" db:"transaction_date"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
