package entity

import "github.com/google/uuid"

type Product struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name" validate:"required"`
	PriceVarejo   float64   `json:"price_varejo" validate:"required"`
	PriceAtacado  float64   `json:"price_atacado" validate:"required"`
	AtacadoAmount int       `json:"atacado_min_amount" validate:"min=1"`
}

func (p Product) GetUnitPrice(amount int) float64 {
	if amount >= p.AtacadoAmount {
		return p.PriceAtacado
	}

	return p.PriceVarejo
}
