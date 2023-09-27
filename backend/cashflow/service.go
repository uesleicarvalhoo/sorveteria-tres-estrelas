package cashflow

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/sales"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/transaction"
)

type Service struct {
	sales        sales.UseCase
	transactions transaction.UseCase
}

func NewService(salesSvc sales.UseCase, transactionSvc transaction.UseCase) Service {
	return Service{
		sales:        salesSvc,
		transactions: transactionSvc,
	}
}

func (s Service) GetCashFlow(ctx context.Context) (CashFlow, error) {
	sales, err := s.sales.GetAll(ctx)
	if err != nil {
		return CashFlow{}, err
	}

	transactions, err := s.transactions.GetTransactions(ctx)
	if err != nil {
		return CashFlow{}, err
	}

	return s.parseCashFlow(sales, transactions), nil
}

func (s Service) GetCashFlowBetween(ctx context.Context, startAt, endAt time.Time) (CashFlow, error) {
	sales, err := s.sales.GetByPeriod(ctx, startAt, endAt)
	if err != nil {
		return CashFlow{}, err
	}

	transactions, err := s.transactions.GetByPeriod(ctx, startAt, endAt)
	if err != nil {
		return CashFlow{}, err
	}

	return s.parseCashFlow(sales, transactions), nil
}

func (s Service) parseCashFlow(sales []sales.Sale, transactions []transaction.Transaction) CashFlow {
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

	for _, t := range transactions {
		var operation BalanceType

		switch t.Type {
		case transaction.Credit:
			totalSales += t.Value
			operation = SaleBalance
		case transaction.Debit:
			totalPayments += t.Value
			operation = PaymentBalance
		}

		details = append(details, Detail{
			Description: t.Description,
			Value:       t.Value,
			Type:        operation,
			Date:        t.CreatedAt,
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
