package sales

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/validator"
)

type Item struct {
	Name      string  `json:"name"`
	UnitPrice float64 `json:"unit_price"`
	Amount    int     `json:"amount"`
}

type Sale struct {
	ID          uuid.UUID   `json:"id"`
	PaymentType PaymentType `json:"payment_type"`
	Items       []Item      `json:"items"`
	Total       float64     `json:"total"`
	Description string      `json:"description"`
	Date        time.Time   `json:"date"`
}

func NewSale(payment PaymentType, description string, total float64, items []Item) (Sale, error) {
	s := Sale{
		ID:          uuid.New(),
		PaymentType: payment,
		Total:       total,
		Description: description,
		Date:        time.Now(),
		Items:       items,
	}

	if err := s.Validate(); err != nil {
		return Sale{}, err
	}

	return s, nil
}

func (s Sale) Validate() error {
	v := validator.New()

	if len(s.Items) == 0 {
		v.AddError("items", "a quantidade mínima de items é 1")
	}

	return v.Validate()
}

func (s Sale) ItemsDescription() string {
	items := []string{}

	for _, item := range s.Items {
		items = append(items, fmt.Sprintf("%s (%d)", item.Name, item.Amount))
	}

	return strings.Join(items, ", ")
}
