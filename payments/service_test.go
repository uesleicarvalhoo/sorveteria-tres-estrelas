package payments_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payments"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/payments/mocks"
)

func TestServiceRegisterPayment(t *testing.T) {
	t.Parallel()

	t.Run("check errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about           string
			value           float32
			desc            string
			mockError       error
			expectedError   error
			expectedPayment payments.Payment
		}{
			{
				about:           "when repository returns an error",
				value:           1,
				desc:            "test repository error",
				mockError:       errors.New("repository error"),
				expectedError:   errors.New("repository error"),
				expectedPayment: payments.Payment{},
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()
				// Arrange
				repo := mocks.NewRepository(t)
				repo.On("Create", mock.Anything, mock.AnythingOfType("payments.Payment")).Return(tc.mockError).Once()

				sut := payments.NewService(repo)

				// Action
				_, err := sut.RegisterPayment(context.Background(), tc.value, tc.desc)

				// Assert
				assert.Equal(t, tc.expectedError, err)
			})
		}
	})

	t.Run("check success", func(t *testing.T) {
		t.Parallel()
		// Arrange
		value := float32(1)
		desc := "test success"

		repo := mocks.NewRepository(t)
		repo.On("Create", mock.Anything, mock.AnythingOfType("payments.Payment")).Return(nil).Once()

		sut := payments.NewService(repo)

		// Action
		p, err := sut.RegisterPayment(context.Background(), value, desc)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, payments.Payment{}, p)
	})
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	repoErr := errors.New("repository error")

	tests := []struct {
		about            string
		mockPayments     []payments.Payment
		mockError        error
		expectedError    error
		expectedPayments []payments.Payment
	}{
		{
			about:         "when repository returns an error",
			mockPayments:  []payments.Payment{},
			mockError:     repoErr,
			expectedError: repoErr,
		},
		{
			about: "when repository returns payments, should order by date",
			mockPayments: []payments.Payment{
				{
					ID:          uuid.Nil,
					Value:       1,
					Description: "test payment 1",
					CreatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          uuid.Nil,
					Value:       2,
					Description: "test payment 2",
					CreatedAt:   time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			mockError:     nil,
			expectedError: nil,
			expectedPayments: []payments.Payment{
				{
					ID:          uuid.Nil,
					Value:       2,
					Description: "test payment 2",
					CreatedAt:   time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          uuid.Nil,
					Value:       1,
					Description: "test payment 1",
					CreatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("GetAll", mock.Anything).Return(tc.mockPayments, tc.mockError).Once()

			sut := payments.NewService(repo)

			// Action
			found, err := sut.GetAll(context.Background())

			// Assert
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedPayments, found)
		})
	}
}

func TestGetBetween(t *testing.T) {
	t.Parallel()

	repoErr := errors.New("repository error")
	tests := []struct {
		about            string
		startAt          time.Time
		endAt            time.Time
		mockPayments     []payments.Payment
		mockError        error
		expectedError    error
		expectedPayments []payments.Payment
	}{
		{
			about:         "when repository returns an error",
			mockPayments:  []payments.Payment{},
			mockError:     repoErr,
			expectedError: repoErr,
		},
		{
			about: "when repository returns payments, should order by date",
			mockPayments: []payments.Payment{
				{
					ID:          uuid.Nil,
					Value:       1,
					Description: "test payment 1",
					CreatedAt:   time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          uuid.Nil,
					Value:       2,
					Description: "test payment 2",

					CreatedAt: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			mockError:     nil,
			expectedError: nil,
			expectedPayments: []payments.Payment{
				{
					ID:          uuid.Nil,
					Value:       2,
					Description: "test payment 2",

					CreatedAt: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          uuid.Nil,
					Value:       1,
					Description: "test payment 1",
					CreatedAt:   time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			startAt: time.Now().Add(-1 * time.Hour),
			endAt:   time.Now().Add(1 * time.Hour),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("GetBetween", mock.Anything, tc.startAt, tc.endAt).Return(tc.mockPayments, tc.mockError).Once()

			sut := payments.NewService(repo)

			// Action
			found, err := sut.GetByPeriod(context.Background(), tc.startAt, tc.endAt)

			// Assert
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedPayments, found)
		})
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		about         string
		id            uuid.UUID
		repoError     error
		expectedError error
	}{
		{
			about:         "when repository returns an error",
			id:            uuid.New(),
			repoError:     errors.New("repository error"),
			expectedError: errors.New("repository error"),
		},
		{
			about:         "when repository returns no error",
			id:            uuid.New(),
			repoError:     nil,
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			repo := mocks.NewRepository(t)
			repo.On("Delete", mock.Anything, tc.id).Return(tc.repoError).Once()

			sut := payments.NewService(repo)

			// Action
			err := sut.DeletePayment(context.Background(), tc.id)

			// Assert
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
