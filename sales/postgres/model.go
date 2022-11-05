package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

type SaleItemModel struct {
	ID     uuid.UUID
	SaleID string
	Name   string
	Price  float64
	Amount int
}

func (SaleItemModel) TableName() string { return "sale_items" }

type SaleModel struct {
	ID          uuid.UUID
	PaymentType sales.PaymentType
	Items       []SaleItemModel `gorm:"foreignKey:SaleID;constraint:OnDelete:CASCADE"`
	Total       float64
	Description string
	Date        time.Time
}

func (SaleModel) TableName() string { return "sales" }

func saleToEntity(m SaleModel) sales.Sale {
	items := make([]sales.Item, len(m.Items))
	for i, item := range m.Items {
		items[i] = sales.Item{
			Name:      item.Name,
			UnitPrice: item.Price,
			Amount:    item.Amount,
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

func saleToModel(s sales.Sale) SaleModel {
	items := make([]SaleItemModel, len(s.Items))

	for i, item := range s.Items {
		items[i] = SaleItemModel{
			ID:     uuid.New(),
			SaleID: s.ID.String(),
			Name:   item.Name,
			Price:  item.UnitPrice,
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
