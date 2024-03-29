package cashflow

import (
	"time"
)

type BalanceType string

const (
	PaymentBalance BalanceType = "Pagamento"
	SaleBalance    BalanceType = "Venda"
)

type Detail struct {
	Description string      `json:"description"`
	Value       float32     `json:"value"`
	Date        time.Time   `json:"date"`
	Type        BalanceType `json:"type"`
}

type CashFlow struct {
	Balance       float32  `json:"balance"`
	TotalSales    float32  `json:"total_sales"`
	TotalPayments float32  `json:"total_payments"`
	Details       []Detail `json:"details"`
}
