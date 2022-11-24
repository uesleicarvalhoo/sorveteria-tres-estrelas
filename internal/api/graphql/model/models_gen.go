// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Balance struct {
	ID          string  `json:"Id"`
	Operation   string  `json:"Operation"`
	Description string  `json:"Description"`
	CreatedAt   string  `json:"CreatedAt"`
	Value       float64 `json:"Value"`
}

type CartItem struct {
	ItemID string `json:"ItemID"`
	Amount int    `json:"Amount"`
}

type CashFlow struct {
	Balances []*Balance `json:"Balances"`
	Payments float64    `json:"Payments"`
	Sales    float64    `json:"Sales"`
	Total    float64    `json:"Total"`
}

type NewBalance struct {
	Description string  `json:"Description"`
	Operation   string  `json:"Operation"`
	Value       float64 `json:"Value"`
}

type NewProduct struct {
	Name          string  `json:"Name"`
	PriceVarejo   float64 `json:"PriceVarejo"`
	PriceAtacado  float64 `json:"PriceAtacado"`
	AtacadoAmount int     `json:"AtacadoAmount"`
}

type NewSale struct {
	Description string      `json:"Description"`
	PaymentType string      `json:"PaymentType"`
	Items       []*CartItem `json:"Items"`
}

type NewUser struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Product struct {
	ID            string  `json:"ID"`
	Name          string  `json:"Name"`
	PriceVarejo   float64 `json:"PriceVarejo"`
	PriceAtacado  float64 `json:"PriceAtacado"`
	AtacadoAmount int     `json:"AtacadoAmount"`
}

type Sale struct {
	ID          string      `json:"id"`
	PaymentType string      `json:"PaymentType"`
	Items       []*SaleItem `json:"Items"`
	Total       float64     `json:"Total"`
	Description string      `json:"Description"`
	Date        string      `json:"Date"`
}

type SaleItem struct {
	Name      string  `json:"Name"`
	UnitPrice float64 `json:"UnitPrice"`
	Amount    int     `json:"Amount"`
}

type SalesByPeriodQuery struct {
	StartAt string `json:"StartAt"`
	EndAt   string `json:"EndAt"`
}

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}