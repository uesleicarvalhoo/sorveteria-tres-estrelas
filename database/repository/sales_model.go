package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

type SaleItemModel struct {
	ID     uuid.UUID
	SaleID uuid.UUID
	Name   string
	Price  float64
	Amount int
}

func (SaleItemModel) TableName() string { return "sale_items" }

type SaleModel struct {
	ID          uuid.UUID
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

	return entity.Sale{
		ID:          entity.ParseUUID(m.ID),
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
			ID:     uuid.New(),
			SaleID: s.ID.ToUUID(),
			Name:   item.Name,
			Price:  item.Price,
			Amount: item.Amount,
		}
	}

	return SaleModel{
		ID:          s.ID.ToUUID(),
		PaymentType: s.PaymentType,
		Total:       s.Total,
		Description: s.Description,
		Date:        s.Date,
		Items:       items,
	}
}
