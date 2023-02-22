package cashflow

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payment"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

type Service struct {
	payments payment.UseCase
	sales    sales.UseCase
}

func NewService(paymentSvc payment.UseCase, salesSvc sales.UseCase) Service {
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

func (s Service) parseCashFlow(payments []payment.Payment, sales []sales.Sale) CashFlow {
	var totalSales, totalPayments float32

	details := []Detail{}

	for _, sale := range sales {
		totalSales += float32(sale.Total)

		details = append(details, Detail{
			Type:        SaleBalance,
			Description: fmt.Sprintf("%s\n%s", sale.Description, sale.ItemsDescription()),
			Value:       float32(sale.Total),
			Date:        sale.Date,
		})
	}

	for _, p := range payments {
		totalPayments += p.Value

		details = append(details, Detail{
			Description: p.Description,
			Value:       p.Value,
			Type:        BalancePayment,
			Date:        p.CreatedAt,
		})
	}

	sort.Slice(details, func(i, j int) bool {
		return details[i].Date.After(details[j].Date)
	})

	return CashFlow{
		Balance:       totalSales - totalPayments,
		TotalSales:    totalSales,
		TotalPayments: totalPayments,
		Details:       details,
	}
}
