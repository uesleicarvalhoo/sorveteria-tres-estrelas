package balances_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
)

func TestBalance(t *testing.T) {
	t.Parallel()

	t.Run("check errors", func(t *testing.T) {
		tests := []struct {
			about         string
			value         float32
			desc          string
			op            balances.OperationType
			expectedError string
		}{
			{
				about:         "when value is zero",
				desc:          "test zero value",
				op:            balances.OperationSale,
				expectedError: balances.ErrInvalidValue.Error(),
			},
			{
				about:         "when description is empty",
				value:         1,
				desc:          "",
				op:            balances.OperationSale,
				expectedError: balances.ErrInvalidDescription.Error(),
			},
			{
				about:         "when operation is invalid",
				value:         1,
				desc:          "test invalid operation",
				op:            balances.OperationType("invalid"),
				expectedError: balances.ErrInvalidOperation.Error(),
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()
				// Arrange

				// Action
				b, err := balances.NewBalance(tc.value, tc.desc, tc.op)

				// Assert
				assert.Equal(t, balances.Balance{}, b)
				assert.EqualError(t, err, tc.expectedError)
			})
		}
	})

	t.Run("check success", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about string
			value float32
			desc  string
			op    balances.OperationType
		}{
			{
				about: "sale operation",
				value: 100,
				desc:  "sale operation test",
				op:    balances.OperationSale,
			},
			{
				about: "payment operation",
				value: 100,
				desc:  "payment operation test",
				op:    balances.OperationPayment,
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()

				b, err := balances.NewBalance(tc.value, tc.desc, tc.op)

				// Assert
				assert.NoError(t, err)

				assert.NotEqual(t, uuid.Nil, b.ID)
				assert.Equal(t, tc.value, b.Value)
				assert.Equal(t, tc.desc, b.Description)
				assert.Equal(t, tc.op, b.Operation)
			})
		}
	})
}
