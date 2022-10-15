package repository

import (
	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

type PopsicleModel struct {
	ID     uuid.UUID
	Flavor string
	Price  float64
}

func (u PopsicleModel) TableName() string { return "users" }

func popsiclelToModel(p entity.Popsicle) PopsicleModel {
	return PopsicleModel{
		ID:     p.ID.ToUUID(),
		Flavor: p.Flavor,
		Price:  p.Price,
	}
}

func popsicleModelToEntity(p PopsicleModel) entity.Popsicle {
	return entity.Popsicle{
		ID:     entity.ParseUUID(p.ID),
		Flavor: p.Flavor,
		Price:  p.Price,
	}
}
