package transaction

import "github.com/jmoiron/sqlx"

type repositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) GetTransactionByID(id uint) (*Transaction, error) {
	var transaction Transaction
	err := r.db.Get(&transaction, "SELECT id, user_id, amount, type, description, transaction_date, created_at, updated_at FROM transactions WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *repositoryImpl) CreateTransaction(transaction *Transaction) (*Transaction, error) {
	query := `
		INSERT INTO transactions (user_id, amount, type, description, transaction_date) 
		VALUES (:user_id, :amount, :type, :description, :transaction_date)
		RETURNING id, user_id, amount, type, description, transaction_date, created_at, updated_at
	`

	rows, err := r.db.NamedQuery(query, transaction)
	if err != nil {
		return nil, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	if !rows.Next() {
		return nil, err
	}

	var createdTransaction Transaction
	if err = rows.StructScan(&createdTransaction); err != nil {
		return nil, err
	}

	return &createdTransaction, nil
}

func (r *repositoryImpl) GetTransactionsByUserID(userID uint) ([]*Transaction, error) {
	var transactions []*Transaction
	err := r.db.Select(&transactions, "SELECT id, user_id, amount, type, description, transaction_date, created_at, updated_at FROM transactions WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
