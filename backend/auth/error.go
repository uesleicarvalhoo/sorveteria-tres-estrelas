package auth

import "errors"

var (
	ErrNotAuthorized = errors.New("usuário ou senha invalidos")
	ErrTokenNotFound = errors.New("token invalido")
	ErrInvalidToken  = errors.New("token invalido")
)
