package cashflow

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payments"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

type CashFlow struct {
	Balance       float32            `json:"balance"`
	TotalSales    float32            `json:"total_sales"`
	TotalPayments float32            `json:"total_payments"`
	Sales         []sales.Sale       `json:"sales"`
	Payments      []payments.Payment `json:"payments"`
}
