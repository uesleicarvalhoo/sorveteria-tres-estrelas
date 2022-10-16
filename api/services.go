package main

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cache"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/database/repository"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/popsicle"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/sales"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/user"
	"gorm.io/gorm"
)

type Services struct {
	authSvc     auth.UseCase
	userSvc     user.UseCase
	popsicleSvc popsicle.UseCase
	salesSvc    sales.UseCase
}

func createServices(db *gorm.DB, cache cache.Cache, secretKey string) *Services {
	popsicleRepo := repository.NewPopsiclePostgres(db)
	salesRepo := repository.NewSalesPostgres(db)
	userRepo := repository.NewUserPostgres(db)

	popsicleSvc := popsicle.NewService(popsicleRepo)
	salesSvc := sales.NewService(popsicleRepo, salesRepo)
	userSvc := user.NewService(userRepo)

	authSvc := auth.NewService(secretKey, userSvc, cache)

	return &Services{
		popsicleSvc: popsicleSvc,
		salesSvc:    salesSvc,
		userSvc:     userSvc,
		authSvc:     authSvc,
	}
}
