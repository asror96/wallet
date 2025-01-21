package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"wallet"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) Send(transaction wallet.Transaction) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (wallet_from, wallet_to, sum) values ($1, $2, $3) RETURNING id", transactionTable)
	row := r.db.QueryRow(query, transaction.From, transaction.To, transaction.Sum)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TransactionPostgres) GetWallet(uid string) (wallet.Wallet, error) {
	var wall wallet.Wallet
	query := fmt.Sprintf("SELECT * FROM %s WHERE uid = $1", walletTable)
	err := r.db.Get(&wall, query, uid)

	return wall, err
}

func (r *TransactionPostgres) UpdateWalletBalance(uid string, newBalance float64) error {
	query := "UPDATE wallets SET balance = $1 WHERE uid = $2"
	_, err := r.db.Exec(query, newBalance, uid)
	return err
}

func (r *TransactionPostgres) ListTransactions(count int64) ([]wallet.Transactions, error) {
	var transactions []wallet.Transactions
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC LIMIT $1", transactionTable)
	err := r.db.Select(&transactions, query, count)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
