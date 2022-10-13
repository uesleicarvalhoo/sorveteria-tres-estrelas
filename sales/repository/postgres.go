package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r PostgresRepository) Get(ctx context.Context, id uuid.UUID) (sales.Sale, error) {
	var m SaleModel

	if tx := r.db.WithContext(ctx).Preload("Items").First(&m); tx.Error != nil {
		return sales.Sale{}, tx.Error
	}

	return toEntity(m), nil
}

func (r PostgresRepository) GetAll(ctx context.Context) ([]sales.Sale, error) {
	var models []SaleModel

	if tx := r.db.WithContext(ctx).Preload("Items").Find(&models); tx.Error != nil {
		return nil, tx.Error
	}

	sales := make([]sales.Sale, len(models))
	for i, m := range models {
		sales[i] = toEntity(m)
	}

	return sales, nil
}

func (r PostgresRepository) Create(ctx context.Context, s sales.Sale) error {
	m := toModel(s)

	return r.db.WithContext(ctx).Create(&m).Error
}

func (r PostgresRepository) Update(ctx context.Context, s sales.Sale) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		m := toModel(s)

		if err := tx.Delete(&SaleItemModel{}, "sale_id = ?", s.ID).Error; err != nil {
			return err
		}

		return tx.Save(&m).Error
	})
}

func (r PostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&SaleModel{}, "id = ?", id).Error
}

func (r PostgresRepository) Search(ctx context.Context, start, end time.Time) ([]sales.Sale, error) {
	var models []SaleModel

	if tx := r.db.WithContext(ctx).Preload("Items").Find(&models, "date BETWEEN ? AND ?", start, end); tx.Error != nil {
		return nil, tx.Error
	}

	sales := make([]sales.Sale, len(models))
	for i, m := range models {
		sales[i] = toEntity(m)
	}

	return sales, nil
}
