package entity

import (
	"errors"

	"github.com/google/uuid"
)

type PaymentType string

const (
	PixPayment       PaymentType = "Pagamento no PIX"
	CashPayment      PaymentType = "Pagamento no Dinheiro"
	CardPayment      PaymentType = "Pagamento no Cartão"
	EmployersPayment PaymentType = "Pagamento de funcionarios"
	SuppliersPayment PaymentType = "Pagamento de fornecedores"
	AnotherPayments  PaymentType = "Outros pagamentos"
)

var ErrInvalidID = errors.New("invalid id")

type ID uuid.UUID

func (id ID) String() string {
	return uuid.UUID(id).String()
}

func (id ID) ToUUID() uuid.UUID {
	return uuid.UUID(id)
}

func StringToID(id string) (ID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return ID{}, ErrInvalidID
	}

	return ID(uid), nil
}

func NewID() ID {
	return ID(uuid.New())
}

func ParseUUID(id uuid.UUID) ID {
	return ID(id)
}
