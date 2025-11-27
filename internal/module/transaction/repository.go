package transaction

import "time"

type Repository interface {
	GetTransactionByID(id uint) (*Transaction, error)
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	GetTransactionsByUserID(userID, limit, offset uint, startDate, endDate *time.Time) ([]*Transaction, error)
	GetAllTransactionsByUserID(userID uint) ([]*Transaction, error)
	GetTransactionSummaryByUserID(userID uint) (*TransactionSummary, error)
}
