package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/products"
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

func (r ProductsPostgres) Get(ctx context.Context, id uuid.UUID) (products.Product, error) {
	var p products.Product

	if tx := r.db.WithContext(ctx).First(&p, "id = ?", id); tx.Error != nil {
		return products.Product{}, tx.Error
	}

	return p, nil
}

func (r ProductsPostgres) GetAll(ctx context.Context) ([]products.Product, error) {
	var records []products.Product

	if tx := r.db.WithContext(ctx).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	return records, nil
}

func (r ProductsPostgres) Create(ctx context.Context, p products.Product) error {
	return r.db.WithContext(ctx).Create(&p).Error
}

func (r ProductsPostgres) Update(ctx context.Context, p *products.Product) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r ProductsPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&products.Product{}, "id = ?", id).Error
}
