package cashflow_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cashflow"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payments"
	mockPayments "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payments/mocks"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
	mockSales "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales/mocks"
)

func TestGetAll(t *testing.T) {
	t.Parallel()

	svcError := errors.New("service error")

	tests := []struct {
		about            string
		paymentsError    error
		storedPayments   []payments.Payment
		salesError       error
		storedSales      []sales.Sale
		expectedCashFlow cashflow.CashFlow
		expectedError    error
	}{
		{
			about:         "when sale service return an error",
			salesError:    svcError,
			expectedError: svcError,
		},
		{
			about:         "when payment service return an error",
			paymentsError: svcError,
			expectedError: svcError,
		},
		{
			about: "should return a cash flow with all sales and payments",
			storedPayments: []payments.Payment{
				{ID: uuid.Nil, Description: "payment 1", Value: 5},
				{ID: uuid.Nil, Description: "payment 2", Value: 3},
				{ID: uuid.Nil, Description: "payment 3", Value: 2},
			},
			storedSales: []sales.Sale{
				{ID: uuid.Nil, Description: "sale 1", Total: 5},
				{ID: uuid.Nil, Description: "sale 2", Total: 14},
			},
			expectedCashFlow: cashflow.CashFlow{
				TotalPayments: 10,
				TotalSales:    19,
				Balance:       9,
				Payments: []payments.Payment{
					{ID: uuid.Nil, Description: "payment 1", Value: 5},
					{ID: uuid.Nil, Description: "payment 2", Value: 3},
					{ID: uuid.Nil, Description: "payment 3", Value: 2},
				},
				Sales: []sales.Sale{
					{ID: uuid.Nil, Description: "sale 1", Total: 5},
					{ID: uuid.Nil, Description: "sale 2", Total: 14},
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

			paymentSvc := mockPayments.NewUseCase(t)
			paymentSvc.On("GetAll", mock.Anything).Return(tc.storedPayments, tc.paymentsError).Maybe()

			saleSvc := mockSales.NewUseCase(t)
			saleSvc.On("GetAll", mock.Anything).Return(tc.storedSales, tc.salesError).Maybe()

			sut := cashflow.NewService(paymentSvc, saleSvc)

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

	svcError := errors.New("service error")

	tests := []struct {
		about            string
		startAt          time.Time
		endAt            time.Time
		paymentsError    error
		storedPayments   []payments.Payment
		salesError       error
		storedSales      []sales.Sale
		expectedCashFlow cashflow.CashFlow
		expectedError    error
	}{
		{
			about:         "when sale service return an error",
			salesError:    svcError,
			expectedError: svcError,
		},
		{
			about:         "when payment service return an error",
			paymentsError: svcError,
			expectedError: svcError,
		},
		{
			about: "should return a cash flow dates with all sales and payments",
			storedPayments: []payments.Payment{
				{ID: uuid.Nil, Description: "payment 1", Value: 15},
				{ID: uuid.Nil, Description: "payment 2", Value: 2},
			},
			storedSales: []sales.Sale{
				{ID: uuid.Nil, Description: "sale 1", Total: 14},
			},
			expectedCashFlow: cashflow.CashFlow{
				TotalPayments: 17,
				TotalSales:    14,
				Balance:       -3,
				Payments: []payments.Payment{
					{ID: uuid.Nil, Description: "payment 1", Value: 15},
					{ID: uuid.Nil, Description: "payment 2", Value: 2},
				},
				Sales: []sales.Sale{
					{ID: uuid.Nil, Description: "sale 1", Total: 14},
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

			paymentSvc := mockPayments.NewUseCase(t)
			paymentSvc.On("GetByPeriod", mock.Anything, tc.startAt, tc.endAt).Return(tc.storedPayments, tc.paymentsError).Maybe()

			saleSvc := mockSales.NewUseCase(t)
			saleSvc.On("GetByPeriod", mock.Anything, tc.startAt, tc.endAt).Return(tc.storedSales, tc.salesError).Maybe()

			sut := cashflow.NewService(paymentSvc, saleSvc)

			// Action
			cf, err := sut.GetCashFlowBetween(context.Background(), tc.startAt, tc.endAt)

			// Assert
			assert.Equal(t, tc.expectedCashFlow, cf)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
