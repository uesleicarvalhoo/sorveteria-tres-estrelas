package entity

import "github.com/google/uuid"

type CartItem struct {
	PopsicleID uuid.UUID `json:"id"`
	Amount     int       `json:"amount"`
}

type Cart struct {
	Items []CartItem `json:"items" validate:"min=1"`
}

func (c *Cart) AddItem(item CartItem) {
	for i := range c.Items {
		if c.Items[i].PopsicleID == item.PopsicleID {
			c.Items[i].Amount += item.Amount

			return
		}
	}

	c.Items = append(c.Items, item)
}
