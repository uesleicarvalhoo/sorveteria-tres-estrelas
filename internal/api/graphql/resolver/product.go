package resolver

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/graphql/model"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/products"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	p, err := r.productSvc.Store(ctx, input.Name, input.PriceVarejo, input.PriceAtacado, input.AtacadoAmount)
	if err != nil {
		return nil, err
	}

	return productFromDomain(p), nil
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (bool, error) {
	pID, err := uuid.Parse(id)
	if err != nil {
		return false, err
	}

	err = r.productSvc.Delete(ctx, pID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	products, err := r.productSvc.Index(ctx)
	if err != nil {
		return nil, err
	}

	productsModel := make([]*model.Product, len(products))
	for i, p := range products {
		productsModel[i] = productFromDomain(p)
	}

	return productsModel, nil
}

func (r *queryResolver) GetProductByID(ctx context.Context, id string) (*model.Product, error) {
	pID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	p, err := r.productSvc.Get(ctx, pID)
	if err != nil {
		return nil, err
	}

	return productFromDomain(p), nil
}

func productFromDomain(p products.Product) *model.Product {
	return &model.Product{
		ID:            p.ID.String(),
		Name:          p.Name,
		PriceVarejo:   p.PriceVarejo,
		PriceAtacado:  p.PriceAtacado,
		AtacadoAmount: p.AtacadoAmount,
	}
}
