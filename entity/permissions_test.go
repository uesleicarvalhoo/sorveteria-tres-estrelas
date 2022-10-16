//go:build unit || all

package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
)

func TestPermission(t *testing.T) {
	t.Parallel()

	tests := []struct {
		about      string
		permission entity.Permission
		domain     string
		actions    []string
	}{
		{
			about:      "sales read write",
			permission: entity.ReadWriteSales,
			domain:     "sales",
			actions:    []string{"read", "write"},
		},
		{
			about:      "sales read",
			permission: entity.ReadSales,
			domain:     "sales",
			actions:    []string{"read"},
		},
		{
			about:      "popsicles read write",
			permission: entity.ReadWritePopsicles,
			domain:     "popsicles",
			actions:    []string{"read", "write"},
		},
		{
			about:      "popsicles read",
			permission: entity.ReadPopsicles,
			domain:     "popsicles",
			actions:    []string{"read"},
		},
		{
			about:      "users read",
			permission: entity.ReadUsers,
			domain:     "users",
			actions:    []string{"read"},
		},
		{
			about:      "users read write",
			permission: entity.ReadWriteUsers,
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
