package auth

import "errors"

var (
	ErrNotAuthorized = errors.New("usuário ou senha invalidos")
	ErrNotPermited   = errors.New("você não tem permissão para executar essa ação")
	ErrTokenNotFound = errors.New("token invalido")
	ErrInvalidToken  = errors.New("token invalido")
)
