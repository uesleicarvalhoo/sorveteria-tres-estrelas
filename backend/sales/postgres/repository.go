package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/sales"
	"gorm.io/gorm"
)

type SalesPostgres struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *SalesPostgres {
	return &SalesPostgres{
		db: db,
	}
}

func (r SalesPostgres) Get(ctx context.Context, id uuid.UUID) (sales.Sale, error) {
	var m SaleModel

	if tx := r.db.WithContext(ctx).Preload("Items").First(&m, "id = ?", id); tx.Error != nil {
		return sales.Sale{}, tx.Error
	}

	return saleToEntity(m), nil
}

func (r SalesPostgres) GetAll(ctx context.Context) ([]sales.Sale, error) {
	var models []SaleModel

	if tx := r.db.WithContext(ctx).Preload("Items").Find(&models); tx.Error != nil {
		return nil, tx.Error
	}

	sales := make([]sales.Sale, len(models))
	for i, m := range models {
		sales[i] = saleToEntity(m)
	}

	return sales, nil
}

func (r SalesPostgres) Create(ctx context.Context, s sales.Sale) error {
	m := saleToModel(s)

	return r.db.WithContext(ctx).Create(&m).Error
}

func (r SalesPostgres) Update(ctx context.Context, s sales.Sale) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		m := saleToModel(s)

		if err := tx.Delete(&SaleItemModel{}, "sale_id = ?", s.ID).Error; err != nil {
			return err
		}

		return tx.Save(&m).Error
	})
}

func (r SalesPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&SaleModel{}, "id = ?", id).Error
}

func (r SalesPostgres) Search(ctx context.Context, start, end time.Time) ([]sales.Sale, error) {
	var models []SaleModel

	if tx := r.db.WithContext(ctx).Preload("Items").Find(&models, "date BETWEEN ? AND ?", start, end); tx.Error != nil {
		return nil, tx.Error
	}

	sales := make([]sales.Sale, len(models))
	for i, m := range models {
		sales[i] = saleToEntity(m)
	}

	return sales, nil
}
