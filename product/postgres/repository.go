package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/product"
	"gorm.io/gorm"
)

type ProductsPostgres struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *ProductsPostgres {
	return &ProductsPostgres{
		db: db,
	}
}

func (r ProductsPostgres) Get(ctx context.Context, id uuid.UUID) (product.Product, error) {
	var p product.Product

	if tx := r.db.WithContext(ctx).First(&p, "id = ?", id); tx.Error != nil {
		return product.Product{}, tx.Error
	}

	return p, nil
}

func (r ProductsPostgres) GetAll(ctx context.Context) ([]product.Product, error) {
	var records []product.Product

	if tx := r.db.WithContext(ctx).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r ProductsPostgres) Create(ctx context.Context, p product.Product) error {
	return r.db.WithContext(ctx).Create(&p).Error
}

func (r ProductsPostgres) Update(ctx context.Context, p *product.Product) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r ProductsPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&product.Product{}, "id = ?", id).Error
}
