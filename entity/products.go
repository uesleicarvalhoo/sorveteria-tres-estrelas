package entity

import (
	"strings"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity/validator"
)

type Product struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name" validate:"required"`
	PriceVarejo   float64   `json:"price_varejo" validate:"required"`
	PriceAtacado  float64   `json:"price_atacado" validate:"required"`
	AtacadoAmount int       `json:"atacado_amount" validate:"min=1"`
}

func NewProduct(name string, priceVarejo, priceAtacado float64, atacadoAmount int) (Product, error) {
	p := Product{
		ID:            uuid.New(),
		Name:          strings.TrimSpace(name),
		PriceVarejo:   priceVarejo,
		PriceAtacado:  priceAtacado,
		AtacadoAmount: atacadoAmount,
	}

	if err := p.Validate(); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (p Product) GetUnitPrice(amount int) float64 {
	if amount >= p.AtacadoAmount {
		return p.PriceAtacado
	}

	return p.PriceVarejo
}

func (p Product) Validate() error {
	v := validator.New()

	if p.Name == "" {
		v.AddError("nome", "campo obrigatório")
	}

	if p.PriceVarejo == 0 {
		v.AddError("preço varejo", "precisa ser maior do que 0")
	}

	if p.PriceAtacado == 0 {
		v.AddError("preço atacado", "precisa ser maior do que 0")
	}

	if p.AtacadoAmount == 0 {
		v.AddError("quantidade atacado", "precisa ser maior do que 0")
	}

	return v.Validate()
}
