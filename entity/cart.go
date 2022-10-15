package entity

type CartItem struct {
	PopsicleID ID  `json:"id"`
	Amount     int `json:"amount"`
}

type Cart struct {
	Items []CartItem `json:"items" validate:"min=1"`
}
