package wallet

type Wallet struct {
	Uid     string  `json:"uid" binding:"required"`
	Balance float64 `json:"balance" binding:"required"`
}
