package product

import "errors"

var ErrNoDataForUpdate = errors.New("no data for update")

type UpdatePayload struct {
	Value         float32 `json:"value"`
	Name          string  `json:"name"`
	PriceVarejo   float64 `json:"price_varejo"`
	PriceAtacado  float64 `json:"price_atacado"`
	AtacadoAmount int     `json:"atacado_amount"`
}

func (up UpdatePayload) IsEmpty() bool {
	return up.Value == 0 && up.Name == "" && up.PriceVarejo == 0 && up.PriceAtacado == 0 && up.AtacadoAmount == 0
}
