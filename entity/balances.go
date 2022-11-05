package entity

import "time"

type Balance struct {
	Value       string
	Date        time.Time
	Description string
}
