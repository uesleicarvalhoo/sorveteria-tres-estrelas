package sales

import (
	"time"

	"github.com/google/uuid"
)

type PaymentType string

const (
	PixPayment       PaymentType = "Pagamento no PIX"
	CashPayment      PaymentType = "Pagamento no Dinheiro"
	CardPayment      PaymentType = "Pagamento no Cartão"
	EmployersPayment PaymentType = "Pagamento de funcionarios"
	SuppliersPayment PaymentType = "Pagamento de fornecedores"
	AnotherPayments  PaymentType = "Outros pagamentos"
)

type SaleItem struct {
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Amount int     `json:"amount"`
}

type Sale struct {
	ID          uuid.UUID   `json:"id"`
	PaymentType PaymentType `json:"paymentType"`
	Items       []SaleItem  `json:"items" validate:"required,min=1"`
	Total       float32     `json:"total"`
	Description string      `json:"description"`
	Date        time.Time   `json:"date"`
}

type CartItem struct {
	PopsicleID uuid.UUID `json:"id"`
	Amount     int       `json:"amount"`
}

type Cart struct {
	Items []CartItem `json:"items" validate:"min=1"`
}
