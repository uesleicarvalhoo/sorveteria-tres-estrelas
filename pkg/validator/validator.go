package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

var v = validator.New() //nolint: gochecknoglobals

func Validate(s any) error {
	err := v.Struct(s)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors) //nolint:errorlint
	if !ok {
		return err
	}

	errMessages := make([]string, len(validationErrors))

	for i, err := range validationErrors {
		switch err.Tag() {
		case "required":
			errMessages[i] = fmt.Sprintf("%s é obrigatorio", err.Field())
		case "min":
			errMessages[i] = fmt.Sprintf("A quantidade mínima de %s é %s", err.Field(), err.Param())
		case "email":
			errMessages[i] = fmt.Sprintf("'%s' não é um email valido", err.Value())
		default:
			errMessages[i] = fmt.Sprintf(
				"'%s' tem o valor '%v' que não satisfaz a condição '%s'", err.Field(), err.Value(), err.Tag(),
			)
		}
	}

	return errors.New(strings.Join(errMessages, "\n"))
}
