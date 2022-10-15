package repository

import (
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

type SaleItemModel struct {
	ID     string
	SaleID string
	Name   string
	Price  float64
	Amount int
}

func (SaleItemModel) TableName() string { return "sale_items" }

type SaleModel struct {
	ID          string
	PaymentType entity.PaymentType
	Items       []SaleItemModel `gorm:"foreignKey:SaleID;constraint:OnDelete:CASCADE"`
	Total       float64
	Description string
	Date        time.Time
}

func (SaleModel) TableName() string { return "sales" }

func saleToEntity(m SaleModel) entity.Sale {
	items := make([]entity.SaleItem, len(m.Items))
	for i, item := range m.Items {
		items[i] = entity.SaleItem{
			Name:   item.Name,
			Price:  item.Price,
			Amount: item.Amount,
		}
	}

	id, _ := entity.StringToID(m.ID)

	return entity.Sale{
		ID:          id,
		PaymentType: m.PaymentType,
		Total:       m.Total,
		Description: m.Description,
		Date:        m.Date,
		Items:       items,
	}
}

func saleToModel(s entity.Sale) SaleModel {
	items := make([]SaleItemModel, len(s.Items))

	for i, item := range s.Items {
		items[i] = SaleItemModel{
			ID:     entity.NewID().String(),
			SaleID: s.ID.String(),
			Name:   item.Name,
			Price:  item.Price,
			Amount: item.Amount,
		}
	}

	return SaleModel{
		ID:          s.ID.String(),
		PaymentType: s.PaymentType,
		Total:       s.Total,
		Description: s.Description,
		Date:        s.Date,
		Items:       items,
	}
}
