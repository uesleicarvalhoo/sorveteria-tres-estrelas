package entity

import "github.com/google/uuid"

type Popsicle struct {
	ID     uuid.UUID `json:"id"`
	Flavor string    `json:"flavor" validate:"required,min=4"`
	Price  float64   `json:"price" validate:"required"`
}
