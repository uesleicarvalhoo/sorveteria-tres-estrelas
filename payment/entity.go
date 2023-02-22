package payment

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID          uuid.UUID `json:"id"`
	Value       float32   `json:"value"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
