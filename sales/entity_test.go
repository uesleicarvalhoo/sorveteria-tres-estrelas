package sales_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
)

func TestSaleItemsDescription(t *testing.T) {
	t.Parallel()

	t.Run("should return a string with the items description", func(t *testing.T) {
		t.Parallel()

		items := []sales.Item{
			{Name: "Item 1", Amount: 1},
			{Name: "Item 2", Amount: 2},
		}

		sale := sales.Sale{Items: items}

		got := sale.ItemsDescription()
		want := "Item 1 (1), Item 2 (2)"

		assert.Equal(t, want, got)
	})
}
