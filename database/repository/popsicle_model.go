package repository

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

type PopsicleModel struct {
	ID     string
	Flavor string
	Price  float64
}

func (u PopsicleModel) TableName() string { return "users" }

func popsiclelToModel(p entity.Popsicle) PopsicleModel {
	return PopsicleModel{
		ID:     p.ID.String(),
		Flavor: p.Flavor,
		Price:  p.Price,
	}
}

func popsicleModelToEntity(p PopsicleModel) entity.Popsicle {
	id, _ := entity.StringToID(p.ID)

	return entity.Popsicle{
		ID:     id,
		Flavor: p.Flavor,
		Price:  p.Price,
	}
}
