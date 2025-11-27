package transaction

import "time"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetTransactionByID(id uint) (*TransactionResponse, error) {
	// TODO: implement this
	return nil, nil
}

func (s *Service) CreateTransaction(transactionRequest *CreateTransactionRequest) (*TransactionResponse, error) {
	transaction := &Transaction{
		UserID:          transactionRequest.UserID,
		Amount:          transactionRequest.Amount,
		Type:            transactionRequest.Type,
		Description:     transactionRequest.Description,
		TransactionDate: time.Unix(transactionRequest.TransactionDate, 0),
	}

	createdTransaction, err := s.repo.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	transactionResponse := &TransactionResponse{
		ID:              &createdTransaction.ID,
		UserID:          &createdTransaction.UserID,
		Amount:          createdTransaction.Amount,
		Type:            createdTransaction.Type,
		Description:     createdTransaction.Description,
		TransactionDate: createdTransaction.TransactionDate.Unix(),
		CreatedAt:       createdTransaction.CreatedAt.Unix(),
		UpdatedAt:       createdTransaction.UpdatedAt.Unix(),
	}

	return transactionResponse, nil
}

func (s *Service) GetTransactionsByUserID(req *GetTransactionsByUserIDRequest) ([]*TransactionLoadMoreResponse, error) {
	transactions, err := s.repo.GetTransactionsByUserID(req.UserID, req.Limit+1, req.Offset, time.Unix(req.StartDate, 0), time.Unix(req.EndDate, 0))
	if err != nil {
		return nil, err
	}

	hasMore := len(transactions) > int(req.Limit)

	if hasMore {
		transactions = transactions[:req.Limit]
	}

	var transactionResponses []*TransactionResponse
	for _, tx := range transactions {
		txResp := &TransactionResponse{
			ID:              &tx.ID,
			UserID:          &tx.UserID,
			Amount:          tx.Amount,
			Type:            tx.Type,
			Description:     tx.Description,
			TransactionDate: tx.TransactionDate.Unix(),
			CreatedAt:       tx.CreatedAt.Unix(),
			UpdatedAt:       tx.UpdatedAt.Unix(),
		}
		transactionResponses = append(transactionResponses, txResp)
	}

	return []*TransactionLoadMoreResponse{
		{
			Transactions: transactionResponses,
			HasMore:      hasMore,
		},
	}, nil
}

func (s *Service) GetTransactionSummaryByUserID(userID uint) (*TransactionSummaryResponse, error) {
	transactionSummary, err := s.repo.GetTransactionSummaryByUserID(userID)
	if err != nil {
		return nil, err
	}

	summary := &TransactionSummaryResponse{
		TotalCredit: transactionSummary.TotalCredit,
		TotalDebit:  transactionSummary.TotalDebit,
		Balance:     transactionSummary.TotalCredit - transactionSummary.TotalDebit,
	}

	return summary, nil
}
