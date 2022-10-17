//go:build unit || all

package sales_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	productsMocks "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/products/mocks"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/sales"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/sales/mocks"
)

func TestNewSale(t *testing.T) {
	t.Parallel()

	t.Run("when all fields are ok", func(t *testing.T) {
		t.Parallel()

		// Arrange
		storedProduct := entity.Product{
			ID:            uuid.New(),
			Name:          "picole de chocolate",
			PriceVarejo:   7.5,
			PriceAtacado:  5,
			AtacadoAmount: 10,
		}

		description := "i'm a sale description"
		paymentType := entity.CashPayment

		cart := entity.Cart{
			Items: []entity.CartItem{
				{
					ItemID: storedProduct.ID,
					Amount: 5,
				},
			},
		}

		prodReader := productsMocks.NewReader(t)
		prodReader.On("Get", mock.Anything, storedProduct.ID).Return(storedProduct, nil).Once()

		salesRepo := mocks.NewRepository(t)
		salesRepo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		sut := sales.NewService(prodReader, salesRepo)
		// Action
		sale, err := sut.RegisterSale(context.Background(), description, paymentType, cart)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, sale.ID)
		assert.Equal(t, description, sale.Description)
		assert.Len(t, sale.Items, len(cart.Items))
		assert.Equal(t, storedProduct.Name, sale.Items[0].Name)
		assert.Equal(t, 37.5, sale.Total)
		assert.False(t, sale.Date.IsZero())
	})

	testErrors := []struct {
		about             string
		cart              entity.Cart
		description       string
		payment           entity.PaymentType
		productRepoReturn entity.Product
		productRepoErr    error
		saleRepoErr       error
		expectedErr       string
	}{
		{
			about:       "when cart items are empty",
			cart:        entity.Cart{Items: []entity.CartItem{}},
			payment:     entity.AnotherPayments,
			description: "i'm a sale description",
			expectedErr: "A quantidade mínima de Items é 1",
		},
		{
			about:       "when product don't exist",
			description: "i'm a sale description",
			payment:     entity.AnotherPayments,
			cart: entity.Cart{
				Items: []entity.CartItem{
					{ItemID: uuid.Nil, Amount: 10},
				},
			},
			productRepoErr: errors.New("record not found"),
			expectedErr:    "record not found",
		},
		{
			about:       "when sale repository return an error when sale is created",
			description: "i'm a sale description",
			payment:     entity.AnotherPayments,
			cart: entity.Cart{
				Items: []entity.CartItem{
					{ItemID: uuid.Nil, Amount: 10},
				},
			},
			saleRepoErr: errors.New("failed to create a new sale"),
			expectedErr: "failed to create a new sale",
		},
	}

	for _, tc := range testErrors {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			productR := productsMocks.NewReader(t)
			productR.On("Get", mock.Anything, mock.Anything).Return(tc.productRepoReturn, tc.productRepoErr).Maybe()

			salesRepo := mocks.NewRepository(t)
			salesRepo.On("Create", mock.Anything, mock.Anything).Return(tc.saleRepoErr).Maybe()

			sut := sales.NewService(productR, salesRepo)

			// Action
			sale, err := sut.RegisterSale(context.Background(), tc.description, tc.payment, tc.cart)

			// Assert
			assert.Equal(t, entity.Sale{}, sale)
			assert.EqualError(t, err, tc.expectedErr)
		})
	}
}
