package cashflow

import (
	"context"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payments"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

type Service struct {
	payments payments.UseCase
	sales    sales.UseCase
}

func NewService(paymentSvc payments.UseCase, salesSvc sales.UseCase) Service {
	return Service{
		payments: paymentSvc,
		sales:    salesSvc,
	}
}

func (s Service) GetCashFlow(ctx context.Context) (CashFlow, error) {
	sales, err := s.sales.GetAll(ctx)
	if err != nil {
		return CashFlow{}, err
	}

	payments, err := s.payments.GetAll(ctx)
	if err != nil {
		return CashFlow{}, err
	}

	return s.parseCashFlow(payments, sales), nil
}

func (s Service) GetCashFlowBetween(ctx context.Context, startAt, endAt time.Time) (CashFlow, error) {
	sales, err := s.sales.GetByPeriod(ctx, startAt, endAt)
	if err != nil {
		return CashFlow{}, err
	}

	payments, err := s.payments.GetByPeriod(ctx, startAt, endAt)
	if err != nil {
		return CashFlow{}, err
	}

	return s.parseCashFlow(payments, sales), nil
}

func (s Service) parseCashFlow(payments []payments.Payment, sales []sales.Sale) CashFlow {
	var totalSales, totalPayments float32

	for _, sale := range sales {
		totalSales += float32(sale.Total)
	}

	for _, payment := range payments {
		totalPayments += payment.Value
	}

	return CashFlow{
		Balance:       totalSales - totalPayments,
		TotalSales:    totalSales,
		TotalPayments: totalPayments,
		Sales:         sales,
		Payments:      payments,
	}
}
