package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/graphql/model"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

// RegisterSale is the resolver for the registerSale field.
func (r *mutationResolver) RegisterSale(ctx context.Context, input model.NewSale) (*model.Sale, error) {
	cart, err := cartFromModel(input.Items)
	if err != nil {
		return nil, err
	}

	s, err := r.salesSvc.RegisterSale(ctx, input.Description, sales.PaymentType(input.PaymentType), cart)
	if err != nil {
		return nil, err
	}

	return saleFromDomain(s), nil
}

// SaleByPeriod is the resolver for the saleByPeriod field.
func (r *queryResolver) SaleByPeriod(ctx context.Context, input model.SalesByPeriodQuery) ([]*model.Sale, error) {
	panic(fmt.Errorf("not implemented: SaleByPeriod - saleByPeriod"))
}

func cartFromModel(items []*model.CartItem) (sales.Cart, error) {
	cart := sales.Cart{}

	for _, i := range items {
		id, err := uuid.Parse(i.ItemID)
		if err != nil {
			return sales.Cart{}, err
		}

		cart.AddItem(sales.CartItem{
			ItemID: id,
			Amount: i.Amount,
		})
	}

	return cart, nil
}

func saleFromDomain(s sales.Sale) *model.Sale {
	items := make([]*model.SaleItem, len(s.Items))

	for i, it := range s.Items {
		items[i] = &model.SaleItem{
			Name:      it.Name,
			UnitPrice: it.UnitPrice,
			Amount:    it.Amount,
		}
	}

	return &model.Sale{
		ID:          s.ID.String(),
		PaymentType: string(s.PaymentType),
		Total:       s.Total,
		Description: s.Description,
		Date:        s.Date.String(),
	}
}
