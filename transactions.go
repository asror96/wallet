package wallet

type Transactions struct {
	Id         int64   `json:"id" db:"id"`            // Сопоставляется с колонкой id в базе данных
	WalletFrom string  `json:"from" db:"wallet_from"` // Сопоставляется с колонкой wallet_from
	WalletTo   string  `json:"to" db:"wallet_to"`     // Сопоставляется с колонкой wallet_to
	Sum        float64 `json:"sum" db:"sum"`          // Сопоставляется с колонкой sum
}
