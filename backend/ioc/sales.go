package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/sales"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/sales/postgres"
	"gorm.io/gorm"
)

func NewSaleService(db *gorm.DB) sales.UseCase {
	r := postgres.NewRepository(db)

	productSvc := NewProductService(db)

	return sales.NewService(productSvc, r)
}
