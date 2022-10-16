package entity

import "strings"

type Permission string

const (
	ReadWriteSales Permission = "sales:read,write"
	ReadSales      Permission = "sales:read"

	ReadWritePopsicles Permission = "popsicles:read,write"
	ReadPopsicles      Permission = "popsicles:read"

	ReadUsers       Permission = "users:read"
	ReadWriteUsers  Permission = "users:read,write"
	AdminPermission Permission = "admin:read,write"
)

func (p Permission) Domain() string {
	d, _ := p.getDomainActions()

	return d
}

func (p Permission) Actions() []string {
	_, a := p.getDomainActions()

	return a
}

func (p Permission) StrActions() string {
	return strings.Join(p.Actions(), ",")
}

func (p Permission) getDomainActions() (string, []string) {
	v := strings.Split(string(p), ":")

	if len(v) == 1 {
		return "", []string{}
	}

	return v[0], strings.Split(v[1], ",")
}

func DefaultPermissions() []Permission {
	return []Permission{
		ReadPopsicles,
		ReadSales,
		ReadUsers,
	}
}