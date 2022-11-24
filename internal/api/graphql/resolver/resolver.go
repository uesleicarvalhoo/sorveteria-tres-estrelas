package resolver

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/graphql/generated"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/products"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

type Resolver struct {
	authSvc    auth.UseCase
	userSvc    users.UseCase
	productSvc products.UseCase
	salesSvc   sales.UseCase
	balanceSvc balances.UseCase
}

func NewResolver(
	authSvc auth.UseCase,
	userSvc users.UseCase,
	productSvc products.UseCase,
	salesSvc sales.UseCase,
	balanceSvc balances.UseCase,
) *Resolver {
	return &Resolver{
		authSvc:    authSvc,
		userSvc:    userSvc,
		productSvc: productSvc,
		salesSvc:   salesSvc,
		balanceSvc: balanceSvc,
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
