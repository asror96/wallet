package service

import (
	"wallet"
	"wallet/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Transactions(count int) ([]wallet.Transactions, error) {
	transactions, err := s.repo.ListTransactions(int64(count))
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
