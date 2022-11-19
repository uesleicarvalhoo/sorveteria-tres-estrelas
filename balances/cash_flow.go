package balances

type CashFlow struct {
	Total    float32   `json:"total"`
	Sales    float32   `json:"sales"`
	Payments float32   `json:"payments"`
	Balances []Balance `json:"balances"`
}
