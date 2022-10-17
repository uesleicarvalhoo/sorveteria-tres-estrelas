package dto

type CreateProductPayload struct {
	Name          string  `json:"name"`
	PriceVarejo   float64 `json:"price_varejo"`
	PriceAtacado  float64 `json:"price_atacado"`
	AtacadoAmount int     `json:"atacado_amount"`
}
