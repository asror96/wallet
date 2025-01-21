package repository

import (
	"github.com/jmoiron/sqlx"
)

type WalletPostgres struct {
	db *sqlx.DB
}

func NewWalletPostgres(db *sqlx.DB) *WalletPostgres {
	return &WalletPostgres{db: db}
}

func (w *WalletPostgres) Balance(uid string) (float64, error) {
	var balance float64
	query := "SELECT balance FROM wallets WHERE uid = $1"
	err := w.db.Get(&balance, query, uid)
	return balance, err
}
