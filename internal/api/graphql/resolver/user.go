package resolver

import (
	"context"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/graphql/model"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	u, err := r.userSvc.Create(ctx, input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	return userFromDomain(u), nil
}

func (r *queryResolver) GetMe(ctx context.Context) (*model.User, error) {
	u, ok := ctx.Value(users.CtxKey{}).(*users.User)
	if !ok {
		return nil, auth.ErrNotAuthorized
	}

	return userFromDomain(*u), nil
}

func userFromDomain(u users.User) *model.User {
	return &model.User{
		ID:           u.ID.String(),
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
}
