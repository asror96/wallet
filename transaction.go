package wallet

type Transaction struct {
	Id   int64   `json:"_"`
	From string  `json:"from" binding:"required"`
	To   string  `json:"to" binding:"required"`
	Sum  float64 `json:"sum" binding:"required"`
}
