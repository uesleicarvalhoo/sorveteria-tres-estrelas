package main

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cache"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/database/repository"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/products"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/sales"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/users"
	"gorm.io/gorm"
)

type Services struct {
	auth     auth.UseCase
	users    users.UseCase
	sales    sales.UseCase
	products products.UseCase
}

func createServices(db *gorm.DB, cache cache.Cache, secretKey string) *Services {
	productsRepo := repository.NewProductsPostgres(db)
	salesRepo := repository.NewSalesPostgres(db)
	userRepo := repository.NewUserPostgres(db)

	userSvc := users.NewService(userRepo)

	return &Services{
		users:    userSvc,
		products: products.NewService(productsRepo),
		sales:    sales.NewService(productsRepo, salesRepo),
		auth:     auth.NewService(secretKey, userSvc, cache),
	}
}
