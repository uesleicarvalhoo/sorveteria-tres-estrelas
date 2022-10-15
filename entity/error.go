package entity

import (
	"fmt"
)

var ErrTooShortPassword = fmt.Errorf("a senha precisa conter ao menos %d caracters", minPasswordLength)
