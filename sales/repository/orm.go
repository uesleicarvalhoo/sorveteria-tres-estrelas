package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

type SaleItemModel struct {
	ID     uuid.UUID
	SaleID uuid.UUID
	Name   string
	Price  float32
	Amount int
}

func (SaleItemModel) TableName() string { return "sale_items" }

type SaleModel struct {
	ID          uuid.UUID
	PaymentType sales.PaymentType
	Items       []SaleItemModel `gorm:"foreignKey:SaleID"`
	Total       float32
	Description string
	Date        time.Time
}

func (SaleModel) TableName() string { return "sales" }

func toEntity(m SaleModel) sales.Sale {
	items := make([]sales.SaleItem, len(m.Items))
	for i, item := range m.Items {
		items[i] = sales.SaleItem{
			Name:   item.Name,
			Price:  item.Price,
			Amount: item.Amount,
		}
	}

	return sales.Sale{
		ID:          m.ID,
		PaymentType: m.PaymentType,
		Total:       m.Total,
		Description: m.Description,
		Date:        m.Date,
		Items:       items,
	}
}

func toModel(s sales.Sale) SaleModel {
	items := make([]SaleItemModel, len(s.Items))

	for i, item := range s.Items {
		items[i] = SaleItemModel{
			ID:     uuid.New(),
			SaleID: s.ID,
			Name:   item.Name,
			Price:  item.Price,
			Amount: item.Amount,
		}
	}

	return SaleModel{
		ID:          s.ID,
		PaymentType: s.PaymentType,
		Total:       s.Total,
		Description: s.Description,
		Date:        s.Date,
		Items:       items,
	}
}
