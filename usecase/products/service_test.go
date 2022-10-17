//go:build unit || all

package products_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/products"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/products/mocks"
)

func TestStore(t *testing.T) {
	t.Parallel()

	t.Run("when all fields are ok", func(t *testing.T) {
		t.Parallel()

		// Arrange
		name := "picole de coco com chocolate"
		varejoPrice := 1.23
		atacadoPrice := 1.0
		atacadoAmount := 5

		repo := mocks.NewRepository(t)
		repo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		sut := products.NewService(repo)

		// Action
		product, err := sut.Store(context.Background(), name, varejoPrice, atacadoPrice, atacadoAmount)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, product)
		assert.Equal(t, product.Name, name)
		assert.Equal(t, product.PriceVarejo, varejoPrice)
		assert.Equal(t, product.PriceAtacado, atacadoPrice)
		assert.Equal(t, product.AtacadoAmount, atacadoAmount)
	})

	testErrors := []struct {
		about         string
		name          string
		priceVarejo   float64
		priceAtacado  float64
		atacadoAmount int
		mockError     error
		expectedErr   string
	}{
		{
			about:         "when repository return an error",
			name:          "picole de amendoin",
			priceVarejo:   1.0,
			priceAtacado:  0.75,
			atacadoAmount: 10,
			expectedErr:   "error on create product",
			mockError:     errors.New("error on create product"),
		},
		{
			about:         "when name is empty",
			name:          "",
			priceVarejo:   1.0,
			priceAtacado:  0.75,
			atacadoAmount: 10,
			expectedErr:   "Name é obrigatorio",
		},
		{
			about:         "when PriceVarejo is invalid",
			name:          "picole de amendoin",
			priceVarejo:   0,
			priceAtacado:  0.75,
			atacadoAmount: 10,
			expectedErr:   "PriceVarejo é obrigatorio",
		},
		{
			about:         "when PriceAtacado is invalid",
			name:          "picole de amendoin",
			priceVarejo:   1.0,
			priceAtacado:  0,
			atacadoAmount: 10,
			expectedErr:   "PriceAtacado é obrigatorio",
		},
		{
			about:         "when atacadoAmount is invalid",
			name:          "picole de amendoin",
			priceVarejo:   1.0,
			priceAtacado:  0.75,
			atacadoAmount: 0,
			expectedErr:   "A quantidade mínima de AtacadoAmount é 1",
		},
	}

	for _, tc := range testErrors {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("Create", mock.Anything, mock.Anything).Return(tc.mockError).Maybe()
			sut := products.NewService(repo)

			// Action
			p, err := sut.Store(context.Background(), tc.name, tc.priceVarejo, tc.priceAtacado, tc.atacadoAmount)

			// Assert
			assert.Equal(t, entity.Product{}, p)
			assert.EqualError(t, err, tc.expectedErr)
		})
	}
}

func TestServiceGet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		describe  string
		err       error
		productID uuid.UUID
		product   entity.Product
	}{
		{
			describe:  "when product is found",
			err:       nil,
			productID: uuid.Nil,
			product: entity.Product{
				ID:            uuid.Nil,
				Name:          "picole de cooco",
				PriceVarejo:   1,
				PriceAtacado:  0.75,
				AtacadoAmount: 10,
			},
		},
		{
			describe:  "when product isn't found",
			err:       fmt.Errorf("err product not found"),
			productID: uuid.Nil,
			product:   entity.Product{},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.describe, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("Get", mock.Anything, tc.productID).Return(tc.product, tc.err).Once()

			sut := products.NewService(repo)

			// Action
			found, err := sut.Get(context.Background(), tc.productID)

			// Assert
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.product, found)
		})
	}
}

func TestServiceGetAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		describe string
		products []entity.Product
	}{
		{
			describe: "when repository is empty",
			products: []entity.Product{},
		},
		{
			describe: "when repository has products",
			products: []entity.Product{
				{ID: uuid.New(), Name: "picole de coco", PriceVarejo: 1, PriceAtacado: 0.75, AtacadoAmount: 10},
				{ID: uuid.New(), Name: "picole de amendoin", PriceVarejo: 1.25, PriceAtacado: 1, AtacadoAmount: 10},
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.describe, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("GetAll", mock.Anything).Return(tc.products, nil).Once()
			service := products.NewService(repo)

			// Action
			found, err := service.Index(context.Background())

			// Assert
			assert.Equal(t, tc.products, found)
			assert.NoError(t, err)
		})
	}
}

func TestServiceUpdate(t *testing.T) {
	t.Parallel()

	t.Run("when update is ok", func(t *testing.T) {
		t.Parallel()

		// Arrange
		p := entity.Product{
			ID:            uuid.New(),
			Name:          "coco com chocolate",
			PriceVarejo:   1.5,
			PriceAtacado:  1.3,
			AtacadoAmount: 20,
		}

		repo := mocks.NewRepository(t)
		repo.On("Update", mock.Anything, &p).Return(nil).Once()

		sut := products.NewService(repo)

		// Action
		p.Name = "coco com goiaba"
		p.PriceVarejo = 1.25
		p.PriceAtacado = 1.0
		p.AtacadoAmount = 30

		err := sut.Update(context.Background(), &p)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("when update return an errror", func(t *testing.T) {
		t.Parallel()

		// Arrange
		p := entity.Product{ID: uuid.New(), Name: "amendoin", PriceVarejo: 1.25, PriceAtacado: 1, AtacadoAmount: 15}

		mockError := errors.New("failed to update product")
		expectedErr := "failed to update product"

		repo := mocks.NewRepository(t)
		repo.On("Update", mock.Anything, &p).Return(mockError).Once()

		sut := products.NewService(repo)

		// Action
		err := sut.Update(context.Background(), &p)

		// Assert
		assert.EqualError(t, err, expectedErr)
	})

	t.Run("when new entity is invalid", func(t *testing.T) {
		t.Parallel()

		// Arrange
		p := entity.Product{
			ID:            uuid.New(),
			Name:          "coco com chocolate",
			PriceVarejo:   1.0,
			PriceAtacado:  1.25,
			AtacadoAmount: 10,
		}

		repo := mocks.NewRepository(t)
		repo.On("Update", mock.Anything, &p).Return(nil).Maybe()

		sut := products.NewService(repo)

		// Action
		p.Name = ""

		err := sut.Update(context.Background(), &p)

		// Assert
		assert.EqualError(t, err, "Name é obrigatorio")
	})
}

func TestServiceDelete(t *testing.T) {
	t.Parallel()

	t.Run("when delete is ok", func(t *testing.T) {
		t.Parallel()

		// Arrange
		p := entity.Product{
			ID:            uuid.New(),
			Name:          "picole de mangaba",
			PriceVarejo:   1.5,
			PriceAtacado:  1.25,
			AtacadoAmount: 5,
		}

		repo := mocks.NewRepository(t)
		repo.On("Delete", mock.Anything, p.ID).Return(nil).Once()

		sut := products.NewService(repo)

		// Action
		err := sut.Delete(context.Background(), p.ID)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("when ID is not found", func(t *testing.T) {
		t.Parallel()

		// Arrange
		errMsg := "record not found"
		p := entity.Product{}

		repo := mocks.NewRepository(t)
		repo.On("Delete", mock.Anything, p.ID).Return(errors.New(errMsg)).Once()

		sut := products.NewService(repo)

		// Action
		err := sut.Delete(context.Background(), p.ID)

		// Assert
		assert.EqualError(t, err, errMsg)
	})
}
