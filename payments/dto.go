package payments

type UpdatePayload struct {
	Value       float32 `json:"value"`
	Description string  `json:"description"`
}
