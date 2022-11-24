package balances_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances/mocks"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

func TestServiceCreate(t *testing.T) {
	t.Parallel()

	t.Run("check errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about           string
			value           float32
			desc            string
			op              balances.OperationType
			mockError       error
			expectedError   error
			expectedBalance balances.Balance
		}{
			{
				about:           "when repository returns an error",
				value:           1,
				desc:            "test repository error",
				op:              balances.OperationSale,
				mockError:       errors.New("repository error"),
				expectedError:   errors.New("repository error"),
				expectedBalance: balances.Balance{},
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()
				// Arrange
				repo := mocks.NewRepository(t)
				repo.On("Create", mock.Anything, mock.AnythingOfType("balances.Balance")).Return(tc.mockError).Once()

				sut := balances.NewService(repo)

				// Action
				_, err := sut.RegisterOperation(context.Background(), tc.value, tc.desc, tc.op)

				// Assert
				assert.Equal(t, tc.expectedError, err)
			})
		}
	})

	t.Run("check success", func(t *testing.T) {
		t.Parallel()
		// Arrange
		value := float32(1)
		desc := "test success"
		op := balances.OperationSale

		repo := mocks.NewRepository(t)
		repo.On("Create", mock.Anything, mock.AnythingOfType("balances.Balance")).Return(nil).Once()

		sut := balances.NewService(repo)

		// Action
		b, err := sut.RegisterOperation(context.Background(), value, desc, op)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, balances.Balance{}, b)
	})
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	repoErr := errors.New("repository error")
	existingBalances := []balances.Balance{
		{
			ID:          uuid.New(),
			Value:       1,
			Description: "test balance 1",
			Operation:   balances.OperationSale,
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Value:       2,
			Description: "test balance 2",
			Operation:   balances.OperationPayment,
			CreatedAt:   time.Now(),
		},
	}

	tests := []struct {
		about            string
		mockBalances     []balances.Balance
		mockError        error
		expectedError    error
		expectedBalances []balances.Balance
	}{
		{
			about:            "when repository returns an error",
			mockBalances:     []balances.Balance{},
			mockError:        repoErr,
			expectedError:    repoErr,
			expectedBalances: []balances.Balance{},
		},
		{
			about:            "when repository returns balances",
			mockBalances:     existingBalances,
			mockError:        nil,
			expectedError:    nil,
			expectedBalances: existingBalances,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("GetAll", mock.Anything).Return(tc.mockBalances, tc.mockError).Once()

			sut := balances.NewService(repo)

			// Action
			found, err := sut.GetAll(context.Background())

			// Assert
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedBalances, found)
		})
	}
}

func TestGetBetween(t *testing.T) {
	t.Parallel()

	repoErr := errors.New("repository error")
	existingBalances := []balances.Balance{
		{
			ID:          uuid.New(),
			Value:       1,
			Description: "test balance 1",
			Operation:   balances.OperationSale,
			CreatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Value:       2,
			Description: "test balance 2",
			Operation:   balances.OperationPayment,
			CreatedAt:   time.Now(),
		},
	}

	cashFlow := balances.CashFlow{
		Total:    -1,
		Sales:    1,
		Payments: 2,
		Balances: existingBalances,
	}

	tests := []struct {
		about            string
		startAt          time.Time
		endAt            time.Time
		mockBalances     []balances.Balance
		mockError        error
		expectedError    error
		expectecCashFlow balances.CashFlow
	}{
		{
			about:            "when repository returns an error",
			mockBalances:     []balances.Balance{},
			mockError:        repoErr,
			expectedError:    repoErr,
			expectecCashFlow: balances.CashFlow{},
		},
		{
			about:            "when repository returns balances",
			mockBalances:     existingBalances,
			mockError:        nil,
			expectedError:    nil,
			expectecCashFlow: cashFlow,
			startAt:          time.Now().Add(-1 * time.Hour),
			endAt:            time.Now().Add(1 * time.Hour),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("GetBetween", mock.Anything, tc.startAt, tc.endAt).Return(tc.mockBalances, tc.mockError).Once()

			sut := balances.NewService(repo)

			// Action
			found, err := sut.GetCashFlowBetween(context.Background(), tc.startAt, tc.endAt)

			// Assert
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectecCashFlow, found)
		})
	}
}

func TestRegisterFromSale(t *testing.T) {
	t.Parallel()

	t.Run("check errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about     string
			sale      sales.Sale
			mockError error
		}{}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()

				// Arrange
				repo := mocks.NewRepository(t)
				repo.On("Create", mock.AnythingOfType("balances.Balance")).Return(tc.mockError).Once()

				sut := balances.NewService(repo)

				// Action
				b, err := sut.RegisterFromSale(context.Background(), tc.sale)

				// Assert
				assert.Equal(t, tc.mockError, err)
				assert.Equal(t, balances.Balance{}, b)
			})
		}
	})

	t.Run("check success", func(t *testing.T) {
		t.Parallel()

		// Arrange
		sale := sales.Sale{
			PaymentType: sales.PixPayment,
			Total:       5,
			Description: "test register balance from sale",
			Date:        time.Now(),
			Items: []sales.Item{
				{Name: "test item 1", UnitPrice: 1, Amount: 1},
				{Name: "test item 2", UnitPrice: 2, Amount: 2},
			},
		}

		expectedDescription := fmt.Sprintf(
			"%s\nItens:\n%dx %s\n%dx %s",
			sale.Description, sale.Items[0].Amount, sale.Items[0].Name, sale.Items[1].Amount, sale.Items[1].Name)

		repo := mocks.NewRepository(t)
		repo.On("Create", mock.Anything, mock.AnythingOfType("balances.Balance")).Return(nil).Once()

		sut := balances.NewService(repo)

		// Action
		b, err := sut.RegisterFromSale(context.Background(), sale)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, b.ID)
		assert.Equal(t, float32(sale.Total), b.Value)
		assert.Equal(t, expectedDescription, b.Description)
		assert.Equal(t, balances.OperationSale, b.Operation)
	})
}
