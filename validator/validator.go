package validator

import (
	"fmt"
	"strings"
)

type ValidationErrorProps struct {
	Context string
	Message string
}

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

type Validator struct {
	errors []ValidationErrorProps
}

func New() Validator {
	return Validator{
		errors: make([]ValidationErrorProps, 0),
	}
}

func (e *Validator) AddError(context, message string) {
	e.errors = append(e.errors, ValidationErrorProps{Context: context, Message: message})
}

func (e *Validator) HasErrors() bool {
	return len(e.errors) > 0
}

func (e Validator) Validate() error {
	if len(e.errors) == 0 {
		return nil
	}

	errGroup := make(map[string][]string, 0)
	for _, err := range e.errors {
		errGroup[err.Context] = append(errGroup[err.Context], err.Message)
	}

	errMsgs := []string{}

	for k, v := range errGroup {
		msg := strings.Join(v, ", ")
		errMsgs = append(errMsgs, fmt.Sprintf("%s: %s", k, msg))
	}

	return &ValidationError{Message: strings.Join(errMsgs, "\n")}
}
