package resolver

import (
	"context"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/graphql/model"
)

func (r *mutationResolver) NewBalance(ctx context.Context, input model.NewBalance) (*model.Balance, error) {
	b, err := r.balanceSvc.RegisterOperation(
		ctx, float32(input.Value), input.Description, balances.OperationType(input.Operation))
	if err != nil {
		return nil, err
	}

	return balanceFromDomain(b), nil
}

func (r *queryResolver) CashFlow(ctx context.Context) (*model.CashFlow, error) {
	c, err := r.balanceSvc.GetCashFlow(ctx)
	if err != nil {
		return nil, err
	}

	cashFlow := &model.CashFlow{
		Total:    float64(c.Total),
		Payments: float64(c.Payments),
		Sales:    float64(c.Sales),
		Balances: make([]*model.Balance, len(c.Balances)),
	}

	for i, b := range c.Balances {
		cashFlow.Balances[i] = balanceFromDomain(b)
	}

	return cashFlow, nil
}

func balanceFromDomain(b balances.Balance) *model.Balance {
	return &model.Balance{
		ID:          b.ID.String(),
		Operation:   b.Operation.String(),
		Description: b.Description,
		CreatedAt:   b.CreatedAt.String(),
		Value:       float64(b.Value),
	}
}
