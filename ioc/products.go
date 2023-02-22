package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/products"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/products/postgres"
	"gorm.io/gorm"
)

func NewProductService(db *gorm.DB) products.UseCase {
	r := postgres.NewRepository(db)

	return products.NewService(r)
}
