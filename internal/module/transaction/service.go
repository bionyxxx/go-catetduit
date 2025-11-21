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

func (s *Service) GetTransactionsByUserID(userID uint) ([]*TransactionResponse, error) {

	transactions, err := s.repo.GetTransactionsByUserID(userID)
	if err != nil {
		return nil, err
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
	return transactionResponses, nil
}
