package entity

type Popsicle struct {
	ID     ID      `json:"id"`
	Flavor string  `json:"flavor" validate:"required,min=4"`
	Price  float64 `json:"price" validate:"required"`
}
