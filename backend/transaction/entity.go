package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Type string

const (
	Credit Type = "credit"
	Debit  Type = "debit"
)

type Transaction struct {
	ID          uuid.UUID `json:"id"`
	Value       float32   `json:"value"`
	Type        Type      `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
