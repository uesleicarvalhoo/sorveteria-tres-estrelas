package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"gorm.io/gorm"
)

type PopsiclePostgres struct {
	db *gorm.DB
}

func NewPopsiclePostgres(db *gorm.DB) *PopsiclePostgres {
	return &PopsiclePostgres{
		db: db,
	}
}

func (r PopsiclePostgres) Get(ctx context.Context, id uuid.UUID) (entity.Popsicle, error) {
	var m PopsicleModel

	if tx := r.db.WithContext(ctx).First(&m); tx.Error != nil {
		return entity.Popsicle{}, tx.Error
	}

	return popsicleModelToEntity(m), nil
}

func (r PopsiclePostgres) GetAll(ctx context.Context) ([]entity.Popsicle, error) {
	var records []PopsicleModel

	if tx := r.db.WithContext(ctx).Find(&records); tx.Error != nil {
		return nil, tx.Error
	}

	popsicles := make([]entity.Popsicle, len(records))
	for i, r := range records {
		popsicles[i] = popsicleModelToEntity(r)
	}

	return popsicles, nil
}

func (r PopsiclePostgres) Create(ctx context.Context, p entity.Popsicle) error {
	m := popsiclelToModel(p)

	return r.db.WithContext(ctx).Create(&m).Error
}

func (r PopsiclePostgres) Update(ctx context.Context, p *entity.Popsicle) error {
	m := popsiclelToModel(*p)

	return r.db.WithContext(ctx).Save(&m).Error
}

func (r PopsiclePostgres) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&PopsicleModel{}, "id = ?", id).Error
}
