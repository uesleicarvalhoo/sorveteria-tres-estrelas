//go:build unit || all

package popsicle_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/popsicle"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/popsicle/mocks"
)

func TestServiceCreate(t *testing.T) {
	t.Parallel()

	t.Run("when all fields are ok", func(t *testing.T) {
		t.Parallel()

		// Arrange
		flavor := "coco com chocolate"
		price := 1.23

		repo := mocks.NewRepository(t)
		repo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		sut := popsicle.NewService(repo)

		// Action
		pop, err := sut.Store(context.Background(), flavor, price)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, pop)
		assert.Equal(t, flavor, pop.Flavor)
		assert.Equal(t, price, pop.Price)
	})

	testErrors := []struct {
		about       string
		flavor      string
		price       float64
		mockError   error
		expectedErr string
	}{
		{
			about:       "when repository return an error",
			flavor:      "amendoin",
			price:       1.0,
			expectedErr: "error on create popsicle",
			mockError:   errors.New("error on create popsicle"),
		},
		{
			about:       "when flavor is empty",
			flavor:      "",
			price:       1.0,
			expectedErr: "Flavor é obrigatorio",
		},
		{
			about:       "when price is 0",
			flavor:      "goiaba",
			price:       0,
			expectedErr: "Price é obrigatorio",
		},
	}

	for _, tc := range testErrors {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("Create", mock.Anything, mock.Anything).Return(tc.mockError).Maybe()
			sut := popsicle.NewService(repo)

			// Action
			p, err := sut.Store(context.Background(), tc.flavor, tc.price)

			// Assert
			assert.Equal(t, entity.Popsicle{}, p)
			assert.EqualError(t, err, tc.expectedErr)
		})
	}
}

func TestServiceGet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		describe   string
		err        error
		popsicleID uuid.UUID
		popscile   entity.Popsicle
	}{
		{
			describe:   "when popsicle is found",
			err:        nil,
			popsicleID: uuid.Nil,
			popscile: entity.Popsicle{
				ID:     uuid.Nil,
				Flavor: "coco",
				Price:  1.0,
			},
		},
		{
			describe:   "when popsicle isn't found",
			err:        fmt.Errorf("err popsicle not found"),
			popsicleID: uuid.Nil,
			popscile:   entity.Popsicle{},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.describe, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("Get", mock.Anything, tc.popsicleID).Return(tc.popscile, tc.err).Once()

			sut := popsicle.NewService(repo)

			// Action
			found, err := sut.Get(context.Background(), tc.popsicleID)

			// Assert
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.popscile, found)
		})
	}
}

func TestServiceGetAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		describe  string
		popsicles []entity.Popsicle
	}{
		{
			describe:  "when repository is empty",
			popsicles: []entity.Popsicle{},
		},
		{
			describe: "when repository has popsicles",
			popsicles: []entity.Popsicle{
				{ID: uuid.New(), Flavor: "goiaba", Price: 2.0},
				{ID: uuid.New(), Flavor: "manga", Price: 1.3},
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.describe, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("GetAll", mock.Anything).Return(tc.popsicles, nil).Once()
			service := popsicle.NewService(repo)

			// Action
			found, err := service.Index(context.Background())

			// Assert
			assert.Equal(t, tc.popsicles, found)
			assert.NoError(t, err)
		})
	}
}

func TestServiceUpdate(t *testing.T) {
	t.Parallel()

	t.Run("when update is ok", func(t *testing.T) {
		t.Parallel()

		// Arrange
		p := entity.Popsicle{
			ID:     uuid.New(),
			Flavor: "coco com chocolate",
			Price:  1.0,
		}

		repo := mocks.NewRepository(t)
		repo.On("Update", mock.Anything, &p).Return(nil).Once()

		sut := popsicle.NewService(repo)

		// Action
		p.Flavor = "coco com goiaba"
		p.Price = 1.25

		err := sut.Update(context.Background(), &p)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("when update return an errror", func(t *testing.T) {
		t.Parallel()

		// Arrange
		p := entity.Popsicle{ID: uuid.New(), Flavor: "amendoin", Price: 1.25}

		mockError := errors.New("failed to update popsicle")
		expectedErr := "failed to update popsicle"

		repo := mocks.NewRepository(t)
		repo.On("Update", mock.Anything, &p).Return(mockError).Once()

		sut := popsicle.NewService(repo)

		// Action
		err := sut.Update(context.Background(), &p)

		// Assert
		assert.EqualError(t, err, expectedErr)
	})

	t.Run("when new entity is invalid", func(t *testing.T) {
		t.Parallel()

		// Arrange
		p := entity.Popsicle{
			ID:     uuid.New(),
			Flavor: "coco com chocolate",
			Price:  1.0,
		}

		repo := mocks.NewRepository(t)
		repo.On("Update", mock.Anything, &p).Return(nil).Maybe()

		sut := popsicle.NewService(repo)

		// Action
		p.Flavor = ""

		err := sut.Update(context.Background(), &p)

		// Assert
		assert.EqualError(t, err, "Flavor é obrigatorio")
	})
}

func TestServiceDelete(t *testing.T) {
	t.Parallel()

	t.Run("when delete is ok", func(t *testing.T) {
		t.Parallel()

		// Arrange
		p := entity.Popsicle{
			ID:     uuid.New(),
			Flavor: "mangaba",
			Price:  0.5,
		}

		repo := mocks.NewRepository(t)
		repo.On("Delete", mock.Anything, p.ID).Return(nil).Once()

		sut := popsicle.NewService(repo)

		// Action
		err := sut.Delete(context.Background(), p.ID)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("when ID is not found", func(t *testing.T) {
		t.Parallel()

		// Arrange
		errMsg := "record not found"
		p := entity.Popsicle{}

		repo := mocks.NewRepository(t)
		repo.On("Delete", mock.Anything, p.ID).Return(errors.New(errMsg)).Once()

		sut := popsicle.NewService(repo)

		// Action
		err := sut.Delete(context.Background(), p.ID)

		// Assert
		assert.EqualError(t, err, errMsg)
	})
}
