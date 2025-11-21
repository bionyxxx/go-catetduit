package transaction

type Repository interface {
	GetTransactionByID(id uint) (*Transaction, error)
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	GetTransactionsByUserID(userID, limit, offset uint) ([]*Transaction, error)
	GetAllTransactionsByUserID(userID uint) ([]*Transaction, error)
	GetTransactionSummaryByUserID(userID uint) (*TransactionSummary, error)
}
