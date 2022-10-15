package entity

type Permission string

const (
	ReadWriteSalesRole = "sales:read,write"
	ReadSalesRole      = "sales:read"

	ReadWritePopsicle = "popsicle:read,write"
	ReadPopsicle      = "popsicle:read"
)

type User struct {
	ID           ID           `json:"id"`
	Name         string       `json:"name" validate:"required"`
	Email        string       `json:"email" validate:"email"`
	PasswordHash string       `json:"-"`
	Permissions  []Permission `json:"roles"`
}
