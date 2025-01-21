package service

import (
	"wallet"
	"wallet/pkg/repository"
)

type Send interface {
	Send(transaction wallet.Transaction) (int, error)
}

type Transaction interface {
	Transactions(count int) ([]wallet.Transactions, error)
}

type Balance interface {
	Balance(uid string) (float64, error)
}

type Service struct {
	Send
	Transaction
	Balance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Send:        NewSendService(repos.Transaction),
		Transaction: NewTransactionService(repos.Transaction),
		Balance:     NewBalanceService(repos.Wallet),
	}
}
