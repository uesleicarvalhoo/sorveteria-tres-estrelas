package payment_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/payment"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/payment/mocks"
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
			expectedPayment payment.Payment
		}{
			{
				about:           "when repository returns an error",
				value:           1,
				desc:            "test repository error",
				mockError:       errors.New("repository error"),
				expectedError:   errors.New("repository error"),
				expectedPayment: payment.Payment{},
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()
				// Arrange
				repo := mocks.NewRepository(t)
				repo.On("Create", mock.Anything, mock.AnythingOfType("payment.Payment")).Return(tc.mockError).Once()

				sut := payment.NewService(repo)

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
		repo.On("Create", mock.Anything, mock.AnythingOfType("payment.Payment")).Return(nil).Once()

		sut := payment.NewService(repo)

		// Action
		p, err := sut.RegisterPayment(context.Background(), value, desc)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, payment.Payment{}, p)
	})
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	repoErr := errors.New("repository error")

	tests := []struct {
		about            string
		mockPayments     []payment.Payment
		mockError        error
		expectedError    error
		expectedPayments []payment.Payment
	}{
		{
			about:         "when repository returns an error",
			mockPayments:  []payment.Payment{},
			mockError:     repoErr,
			expectedError: repoErr,
		},
		{
			about: "when repository returns payments, should order by date",
			mockPayments: []payment.Payment{
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
			expectedPayments: []payment.Payment{
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

			sut := payment.NewService(repo)

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
		mockPayments     []payment.Payment
		mockError        error
		expectedError    error
		expectedPayments []payment.Payment
	}{
		{
			about:         "when repository returns an error",
			mockPayments:  []payment.Payment{},
			mockError:     repoErr,
			expectedError: repoErr,
		},
		{
			about: "when repository returns payments, should order by date",
			mockPayments: []payment.Payment{
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
			expectedPayments: []payment.Payment{
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

			sut := payment.NewService(repo)

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

			sut := payment.NewService(repo)

			// Action
			err := sut.DeletePayment(context.Background(), tc.id)

			// Assert
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	t.Run("check errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about           string
			payment         payment.Payment
			newDescription  string
			newValue        float32
			repoGetError    error
			repoUpdateError error
			expectedError   error
		}{
			{
				about:         "when repository get returns an error",
				payment:       payment.Payment{},
				repoGetError:  errors.New("repository get error"),
				expectedError: errors.New("repository get error"),
			},
			{
				about:           "when repository update returns an error",
				payment:         payment.Payment{},
				repoUpdateError: errors.New("repository update error"),
				expectedError:   errors.New("repository update error"),
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()

				// Arrange
				repo := mocks.NewRepository(t)
				repo.On("Get", mock.Anything, tc.payment.ID).Return(tc.payment, tc.repoGetError).Once()
				repo.On("Update", mock.Anything, mock.Anything).Return(tc.repoUpdateError).Maybe()

				sut := payment.NewService(repo)

				// Action
				p, err := sut.UpdatePayment(context.Background(), tc.payment.ID, tc.newValue, tc.newDescription)

				// Assert
				assert.Equal(t, payment.Payment{}, p)
				assert.Equal(t, tc.expectedError, err)
			})
		}
	})

	t.Run("check success", func(t *testing.T) {
		t.Parallel()

		storedPayment := payment.Payment{}

		// Arrange

		repo := mocks.NewRepository(t)
		repo.On("Get", mock.Anything, storedPayment.ID).Return(storedPayment, nil).Once()
		repo.On("Update", mock.Anything, mock.Anything).Return(nil).Once()

		sut := payment.NewService(repo)

		// Action
		p, err := sut.UpdatePayment(context.Background(), storedPayment.ID, 1, "test")

		// Assert
		assert.Equal(t, p.ID, storedPayment.ID)
		assert.Equal(t, p.Description, "test")
		assert.Equal(t, p.Value, float32(1))
		assert.NoError(t, err)
	})
}
