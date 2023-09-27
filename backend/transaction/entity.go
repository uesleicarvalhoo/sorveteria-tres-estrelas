package transaction

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Type string

const (
	Credit Type = "Venda"
	Debit  Type = "Pagamento"
)

func (t Type) Validate() error {
	switch t {
	case Credit:
		return nil
	case Debit:
		return nil
	default:
		return fmt.Errorf("'%s' is not a valid operation type", t)
	}
}

type Transaction struct {
	ID          uuid.UUID `json:"id"`
	Value       float32   `json:"value"`
	Type        Type      `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
