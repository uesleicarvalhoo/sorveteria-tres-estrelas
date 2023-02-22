package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/product"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/product/postgres"
	"gorm.io/gorm"
)

func NewProductService(db *gorm.DB) product.UseCase {
	r := postgres.NewRepository(db)

	return product.NewService(r)
}
