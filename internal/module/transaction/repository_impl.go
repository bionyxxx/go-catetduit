package transaction

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

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

func (r *repositoryImpl) GetTransactionsByUserID(userID, limit, offset uint, startDate, endDate *time.Time) ([]*Transaction, error) {
	var transactions []*Transaction

	query := "SELECT id, user_id, amount, type, description, transaction_date, created_at, updated_at FROM transactions WHERE user_id=$1"
	args := []interface{}{userID}
	argCount := 1

	if startDate != nil && !startDate.IsZero() {
		argCount++
		query += " AND transaction_date >= $" + fmt.Sprintf("%d", argCount)
		args = append(args, *startDate)
	}

	if endDate != nil && !endDate.IsZero() {
		argCount++
		query += " AND transaction_date <= $" + fmt.Sprintf("%d", argCount)
		args = append(args, *endDate)
	}

	argCount++
	query += " ORDER BY transaction_date DESC LIMIT $" + fmt.Sprintf("%d", argCount)
	args = append(args, limit)

	argCount++
	query += " OFFSET $" + fmt.Sprintf("%d", argCount)
	args = append(args, offset)

	err := r.db.Select(&transactions, query, args...)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
func (r *repositoryImpl) GetAllTransactionsByUserID(userID uint) ([]*Transaction, error) {
	var transactions []*Transaction
	err := r.db.Select(&transactions,
		"SELECT id, user_id, amount, type, description, transaction_date, created_at, updated_at FROM transactions WHERE user_id=$1 ORDER BY transaction_date DESC",
		userID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repositoryImpl) GetTransactionSummaryByUserID(userID uint) (*TransactionSummary, error) {
	var summary TransactionSummary
	query := `
		SELECT 
			COALESCE(SUM(CASE WHEN type = $1 THEN amount ELSE 0 END), 0) AS total_credit,
			COALESCE(SUM(CASE WHEN type = $2 THEN amount ELSE 0 END), 0) AS total_debit
		FROM transactions
		WHERE user_id = $3
	`
	err := r.db.Get(&summary, query, TypeCredit, TypeDebit, userID)
	if err != nil {
		return nil, err
	}
	return &summary, nil
}
