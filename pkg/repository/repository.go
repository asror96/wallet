package repository

import (
	"github.com/jmoiron/sqlx"
	"wallet"
)

type Transaction interface {
	Send(transaction wallet.Transaction) (int, error)
	GetWallet(uid string) (wallet.Wallet, error)
	UpdateWalletBalance(uid string, balance float64) error
	ListTransactions(count int64) ([]wallet.Transactions, error)
}

type Wallet interface {
	Balance(uid string) (float64, error)
}

type Repository struct {
	Transaction
	Wallet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransactionPostgres(db),
		Wallet:      NewWalletPostgres(db),
	}
}
