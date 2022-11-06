package main

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
	balancesPostgres "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances/postgres"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cache"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/products"
	productsPostgres "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/products/postgres"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
	salesPostgres "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales/postgres"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
	usersPostgres "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users/postgres"
	"gorm.io/gorm"
)

type Services struct {
	auth     auth.UseCase
	users    users.UseCase
	sales    sales.UseCase
	balances balances.UseCase
	products products.UseCase
}

func createServices(db *gorm.DB, cache cache.Cache, secretKey string) *Services {
	productsRepo := productsPostgres.NewRepository(db)
	salesRepo := salesPostgres.NewRepository(db)
	userRepo := usersPostgres.NewRepository(db)
	balancesRepo := balancesPostgres.NewRepository(db)

	userSvc := users.NewService(userRepo)

	return &Services{
		users:    userSvc,
		products: products.NewService(productsRepo),
		balances: balances.NewService(balancesRepo),
		sales:    sales.NewService(productsRepo, salesRepo),
		auth:     auth.NewService(secretKey, userSvc, cache),
	}
}
