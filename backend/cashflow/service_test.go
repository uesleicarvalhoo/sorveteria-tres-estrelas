package cashflow_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/cashflow"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/sales"
	mockSales "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/sales/mocks"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/transaction"
	mockTransactions "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/transaction/mocks"
)

func TestGetAll(t *testing.T) {
	t.Parallel()

	makeDate := func(year, month, day int) time.Time {
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}

	svcError := errors.New("service error")

	tests := []struct {
		about              string
		salesError         error
		storedSales        []sales.Sale
		storedTransactions []transaction.Transaction
		transactionError   error
		expectedCashFlow   cashflow.CashFlow
		expectedError      error
	}{
		{
			about:         "when sale service return an error",
			salesError:    svcError,
			expectedError: svcError,
		},
		{
			about:            "when payment service return an error",
			transactionError: svcError,
			expectedError:    svcError,
		},
		{
			about: "should return a cash flow with all sales and payments ordered by date",
			storedTransactions: []transaction.Transaction{
				{ID: uuid.Nil, Description: "payment 1", Type: transaction.Debit, Value: 5, CreatedAt: makeDate(2020, 1, 1)},
				{ID: uuid.Nil, Description: "payment 2", Type: transaction.Debit, Value: 3, CreatedAt: makeDate(2020, 1, 5)},
				{ID: uuid.Nil, Description: "payment 3", Type: transaction.Debit, Value: 2, CreatedAt: makeDate(2020, 1, 7)},
			},
			storedSales: []sales.Sale{
				{ID: uuid.Nil, Description: "sale 1", Total: 5, Date: makeDate(2020, 1, 2)},
				{ID: uuid.Nil, Description: "sale 2", Total: 14, Date: makeDate(2020, 1, 8)},
			},
			expectedCashFlow: cashflow.CashFlow{
				TotalPayments: 10,
				TotalSales:    19,
				Balance:       9,
				Details: []cashflow.Detail{
					{Description: "sale 2\n", Value: 14, Type: cashflow.PaymentBalance, Date: makeDate(2020, 1, 8)},
					{Description: "payment 3", Value: 2, Type: cashflow.SaleBalance, Date: makeDate(2020, 1, 7)},
					{Description: "payment 2", Value: 3, Type: cashflow.SaleBalance, Date: makeDate(2020, 1, 5)},
					{Description: "sale 1\n", Value: 5, Type: cashflow.PaymentBalance, Date: makeDate(2020, 1, 2)},
					{Description: "payment 1", Value: 5, Type: cashflow.SaleBalance, Date: makeDate(2020, 1, 1)},
				},
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			saleSvc := mockSales.NewUseCase(t)
			saleSvc.On("GetAll", mock.Anything).Return(tc.storedSales, tc.salesError).Maybe()

			transactionSvc := mockTransactions.NewUseCase(t)
			transactionSvc.On("GetTransactions", mock.Anything).Return(tc.storedTransactions, tc.transactionError).Maybe()

			sut := cashflow.NewService(saleSvc, transactionSvc)

			// Action
			cf, err := sut.GetCashFlow(context.Background())

			// Assert
			assert.Equal(t, tc.expectedCashFlow, cf)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGetByPeriod(t *testing.T) {
	t.Parallel()

	makeDate := func(year, month, day int) time.Time {
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}

	svcError := errors.New("service error")

	tests := []struct {
		about              string
		startAt            time.Time
		endAt              time.Time
		transactionError   error
		storedTransactions []transaction.Transaction
		salesError         error
		storedSales        []sales.Sale
		expectedCashFlow   cashflow.CashFlow
		expectedError      error
	}{
		{
			about:         "when sale service return an error",
			salesError:    svcError,
			expectedError: svcError,
		},
		{
			about:            "when transaction service return an error",
			transactionError: svcError,
			expectedError:    svcError,
		},
		{
			about: "should return a cash flow dates with all sales and payments",
			storedTransactions: []transaction.Transaction{
				{ID: uuid.Nil, Description: "payment 1", Type: transaction.Debit, Value: 15, CreatedAt: makeDate(2020, 1, 1)},
				{ID: uuid.Nil, Description: "payment 2", Type: transaction.Debit, Value: 2, CreatedAt: makeDate(2020, 1, 5)},
			},
			storedSales: []sales.Sale{
				{ID: uuid.Nil, Description: "sale 1", Total: 14, Date: makeDate(2020, 1, 2)},
			},
			expectedCashFlow: cashflow.CashFlow{
				TotalPayments: 17,
				TotalSales:    14,
				Balance:       -3,
				Details: []cashflow.Detail{
					{Description: "payment 2", Value: 2, Type: cashflow.PaymentBalance, Date: makeDate(2020, 1, 5)},
					{Description: "sale 1\n", Value: 14, Type: cashflow.SaleBalance, Date: makeDate(2020, 1, 2)},
					{Description: "payment 1", Value: 15, Type: cashflow.PaymentBalance, Date: makeDate(2020, 1, 1)},
				},
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			saleSvc := mockSales.NewUseCase(t)
			saleSvc.On("GetByPeriod", mock.Anything, tc.startAt, tc.endAt).Return(tc.storedSales, tc.salesError).Maybe()

			transactionSvc := mockTransactions.NewUseCase(t)
			transactionSvc.On("GetByPeriod", mock.Anything, tc.startAt, tc.endAt).
				Return(tc.storedTransactions, tc.transactionError).Maybe()

			sut := cashflow.NewService(saleSvc, transactionSvc)

			// Action
			cf, err := sut.GetCashFlowBetween(context.Background(), tc.startAt, tc.endAt)

			// Assert
			assert.Equal(t, tc.expectedCashFlow, cf)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
