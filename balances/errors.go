package balances

import "errors"

var (
	ErrInvalidOperation   = errors.New("tipo de operação inválida")
	ErrInvalidValue       = errors.New("valor inválido")
	ErrInvalidDescription = errors.New("descrição inválida")
)
