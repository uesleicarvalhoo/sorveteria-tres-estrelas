package dto

type CreatePopsiclePayload struct {
	Flavor string  `json:"flavor"`
	Price  float64 `json:"price"`
}
