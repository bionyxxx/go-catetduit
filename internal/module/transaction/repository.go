package transaction

type Repository interface {
	GetTransactionByID(id uint) (*Transaction, error)
	CreateTransaction(transaction *Transaction) (*Transaction, error)
	GetTransactionsByUserID(userID uint) ([]*Transaction, error)
}
