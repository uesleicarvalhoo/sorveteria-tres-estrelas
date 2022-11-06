package balances

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID          uuid.UUID     `json:"id"`
	Value       float32       `json:"value"`
	Description string        `json:"description"`
	Operation   OperationType `json:"operation"`
	CreatedAt   time.Time     `json:"created_at"`
}

func NewBalance(value float32, desc string, op OperationType) (Balance, error) {
	b := Balance{
		ID:          uuid.New(),
		Value:       value,
		Description: desc,
		Operation:   op,
		CreatedAt:   time.Now(),
	}

	if err := b.Validate(); err != nil {
		return Balance{}, err
	}

	return b, nil
}

func (b *Balance) Validate() error {
	if b.Value <= 0 {
		return ErrInvalidValue
	}

	if b.Description == "" {
		return ErrInvalidDescription
	}

	if b.Operation != OperationSale && b.Operation != OperationPayment {
		return ErrInvalidOperation
	}

	return nil
}
