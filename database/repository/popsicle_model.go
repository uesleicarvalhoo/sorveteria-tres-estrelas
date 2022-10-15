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

func (u PopsicleModel) TableName() string { return "popsicles" }

func popsiclelToModel(p entity.Popsicle) PopsicleModel {
	return PopsicleModel{
		ID:     p.ID,
		Flavor: p.Flavor,
		Price:  p.Price,
	}
}

func popsicleModelToEntity(p PopsicleModel) entity.Popsicle {
	return entity.Popsicle{
		ID:     p.ID,
		Flavor: p.Flavor,
		Price:  p.Price,
	}
}
