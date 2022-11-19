package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales/postgres"
	"gorm.io/gorm"
)

func NewSaleService(db *gorm.DB) sales.UseCase {
	r := postgres.NewRepository(db)

	productSvc := NewProductService(db)

	return sales.NewService(productSvc, r)
}
