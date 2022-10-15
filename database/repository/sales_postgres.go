package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"gorm.io/gorm"
)

type SalesPostgres struct {
	db *gorm.DB
}

func NewSalesPostgres(db *gorm.DB) *SalesPostgres {
	return &SalesPostgres{
		db: db,
	}
}

func (r SalesPostgres) Get(ctx context.Context, id uuid.UUID) (entity.Sale, error) {
	var m SaleModel

	if tx := r.db.WithContext(ctx).Preload("Items").First(&m, "id = ?", id); tx.Error != nil {
		return entity.Sale{}, tx.Error
	}

	return saleToEntity(m), nil
}

func (r SalesPostgres) GetAll(ctx context.Context) ([]entity.Sale, error) {
	var models []SaleModel

	if tx := r.db.WithContext(ctx).Preload("Items").Find(&models); tx.Error != nil {
		return nil, tx.Error
	}

	sales := make([]entity.Sale, len(models))
	for i, m := range models {
		sales[i] = saleToEntity(m)
	}

	return sales, nil
}

func (r SalesPostgres) Create(ctx context.Context, s entity.Sale) error {
	m := saleToModel(s)

	return r.db.WithContext(ctx).Create(&m).Error
}

func (r SalesPostgres) Update(ctx context.Context, s entity.Sale) error {
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

func (r SalesPostgres) Search(ctx context.Context, start, end time.Time) ([]entity.Sale, error) {
	var models []SaleModel

	if tx := r.db.WithContext(ctx).Preload("Items").Find(&models, "date BETWEEN ? AND ?", start, end); tx.Error != nil {
		return nil, tx.Error
	}

	sales := make([]entity.Sale, len(models))
	for i, m := range models {
		sales[i] = saleToEntity(m)
	}

	return sales, nil
}
