//go:build unit || all

package users_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

func TestPermission(t *testing.T) {
	t.Parallel()

	tests := []struct {
		about      string
		permission users.Permission
		domain     string
		actions    []string
	}{
		{
			about:      "sales read write",
			permission: users.ReadWriteSales,
			domain:     "sales",
			actions:    []string{"read", "write"},
		},
		{
			about:      "sales read",
			permission: users.ReadSales,
			domain:     "sales",
			actions:    []string{"read"},
		},
		{
			about:      "products read write",
			permission: users.ReadWriteProducts,
			domain:     "products",
			actions:    []string{"read", "write"},
		},
		{
			about:      "products read",
			permission: users.ReadProducts,
			domain:     "products",
			actions:    []string{"read"},
		},
		{
			about:      "users read",
			permission: users.ReadUsers,
			domain:     "users",
			actions:    []string{"read"},
		},
		{
			about:      "users read write",
			permission: users.ReadWriteUsers,
			domain:     "users",
			actions:    []string{"read", "write"},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()
			// Action
			domain := tc.permission.Domain()
			actions := tc.permission.Actions()

			// Assert
			assert.Equal(t, tc.domain, domain)
			assert.ElementsMatch(t, tc.actions, actions)
		})
	}
}
