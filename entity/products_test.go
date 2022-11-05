package entity_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

func TestProductGetPrice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		about         string
		product       entity.Product
		amount        int
		expectedPrice float64
	}{
		{
			about:         "when amount is lower then atacadoAmount",
			product:       entity.Product{ID: uuid.New(), Name: "test", PriceVarejo: 1, PriceAtacado: 0.5, AtacadoAmount: 10},
			amount:        9,
			expectedPrice: 1,
		},
		{
			about:         "when amount greather then atacadoAmount",
			product:       entity.Product{ID: uuid.New(), Name: "test", PriceVarejo: 1, PriceAtacado: 0.5, AtacadoAmount: 10},
			amount:        15,
			expectedPrice: 0.5,
		},
		{
			about:         "when amount iqual to atacadoAmount",
			product:       entity.Product{ID: uuid.New(), Name: "test", PriceVarejo: 1, PriceAtacado: 0.5, AtacadoAmount: 10},
			amount:        10,
			expectedPrice: 0.5,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			price := tc.product.GetUnitPrice(tc.amount)

			assert.Equal(t, tc.expectedPrice, price)
		})

	}
}
