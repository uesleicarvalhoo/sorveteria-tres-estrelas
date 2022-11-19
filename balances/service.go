package balances

import (
	"context"
	"fmt"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

type Service struct {
	r Repository
}

func NewService(r Repository) Service {
	return Service{r: r}
}

func (s Service) RegisterOperation(ctx context.Context, value float32, desc string, tp OperationType) (Balance, error) {
	b, err := NewBalance(value, desc, tp)
	if err != nil {
		return Balance{}, err
	}

	if err := s.r.Create(ctx, b); err != nil {
		return Balance{}, err
	}

	return b, nil
}

func (s Service) GetAll(ctx context.Context) ([]Balance, error) {
	return s.r.GetAll(ctx)
}

func (s Service) RegisterFromSale(ctx context.Context, sale sales.Sale) (Balance, error) {
	description := fmt.Sprintf("%s\nItens:", sale.Description)

	for _, item := range sale.Items {
		description += fmt.Sprintf("\n%dx %s", item.Amount, item.Name)
	}

	return s.RegisterOperation(ctx, float32(sale.Total), description, OperationSale)
}

func (s Service) GetCashFlow(ctx context.Context) (CashFlow, error) {
	balances, err := s.r.GetAll(ctx)
	if err != nil {
		return CashFlow{}, err
	}

	var total, sales, payments float32

	for _, b := range balances {
		if b.Operation == OperationSale {
			sales += b.Value
			total += b.Value
		} else {
			payments += b.Value
			total -= b.Value
		}
	}

	return CashFlow{
		Total:    total,
		Payments: payments,
		Sales:    sales,
		Balances: balances,
	}, nil
}
