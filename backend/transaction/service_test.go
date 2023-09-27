package transaction_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/transaction"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/transaction/mocks"
)

func TestServiceRegisterPayment(t *testing.T) {
	t.Parallel()

	t.Run("check errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			about           string
			value           float32
			desc            string
			operation       transaction.Type
			mockError       error
			expectedError   error
			expectedPayment transaction.Transaction
		}{
			{
				about:           "when repository returns an error",
				value:           1,
				desc:            "test repository error",
				operation:       transaction.Credit,
				mockError:       errors.New("repository error"),
				expectedError:   errors.New("repository error"),
				expectedPayment: transaction.Transaction{},
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.about, func(t *testing.T) {
				t.Parallel()
				// Arrange
				repo := mocks.NewRepository(t)
				repo.On("Create", mock.Anything, mock.AnythingOfType("transaction.Transaction")).Return(tc.mockError).Once()

				sut := transaction.NewService(repo)

				// Action
				_, err := sut.RegisterTransaction(context.Background(), tc.value, tc.operation, tc.desc)

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
		operation := transaction.Debit

		repo := mocks.NewRepository(t)
		repo.On("Create", mock.Anything, mock.AnythingOfType("transaction.Transaction")).Return(nil).Once()

		sut := transaction.NewService(repo)

		// Action
		p, err := sut.RegisterTransaction(context.Background(), value, operation, desc)

		// Assert
		assert.NoError(t, err)
		assert.NotEqual(t, transaction.Transaction{}, p)
	})
}

func TestGetAll(t *testing.T) {
	t.Parallel()

	repoErr := errors.New("repository error")

	tests := []struct {
		about                 string
		mockTransactions      []transaction.Transaction
		mockError             error
		expectedError         error
		expectedTransatctions []transaction.Transaction
	}{
		{
			about:            "when repository returns an error",
			mockTransactions: []transaction.Transaction{},
			mockError:        repoErr,
			expectedError:    repoErr,
		},
		{
			about: "when repository returns transactions, should order by date",
			mockTransactions: []transaction.Transaction{
				{
					ID:          uuid.Nil,
					Value:       1,
					Description: "test transaction 1",
					Type:        transaction.Credit,
					CreatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          uuid.Nil,
					Value:       2,
					Type:        transaction.Credit,
					Description: "test transaction 2",
					CreatedAt:   time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			mockError:     nil,
			expectedError: nil,
			expectedTransatctions: []transaction.Transaction{
				{
					ID:          uuid.Nil,
					Value:       2,
					Description: "test transaction 2",
					Type:        transaction.Credit,
					CreatedAt:   time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          uuid.Nil,
					Value:       1,
					Description: "test transaction 1",
					Type:        transaction.Credit,
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
			repo.On("GetAll", mock.Anything).Return(tc.mockTransactions, tc.mockError).Once()

			sut := transaction.NewService(repo)

			// Action
			found, err := sut.GetTransactions(context.Background())

			// Assert
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedTransatctions, found)
		})
	}
}

func TestGetBetween(t *testing.T) {
	t.Parallel()

	repoErr := errors.New("repository error")
	tests := []struct {
		about                string
		startAt              time.Time
		endAt                time.Time
		mockedTransatctions  []transaction.Transaction
		mockError            error
		expectedError        error
		expectedTransactions []transaction.Transaction
	}{
		{
			about:               "when repository returns an error",
			mockedTransatctions: []transaction.Transaction{},
			mockError:           repoErr,
			expectedError:       repoErr,
		},
		{
			about: "when repository returns transactions, should order by date",
			mockedTransatctions: []transaction.Transaction{
				{
					ID:          uuid.Nil,
					Value:       1,
					Type:        transaction.Debit,
					Description: "test transaction 1",
					CreatedAt:   time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          uuid.Nil,
					Value:       2,
					Type:        transaction.Credit,
					Description: "test transaction 2",

					CreatedAt: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			mockError:     nil,
			expectedError: nil,
			expectedTransactions: []transaction.Transaction{
				{
					ID:          uuid.Nil,
					Value:       2,
					Type:        transaction.Credit,
					Description: "test transaction 2",

					CreatedAt: time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          uuid.Nil,
					Value:       1,
					Type:        transaction.Debit,
					Description: "test transaction 1",
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
			repo.On("GetBetween", mock.Anything, tc.startAt, tc.endAt).Return(tc.mockedTransatctions, tc.mockError).Once()

			sut := transaction.NewService(repo)

			// Action
			found, err := sut.GetByPeriod(context.Background(), tc.startAt, tc.endAt)

			// Assert
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedTransactions, found)
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

			sut := transaction.NewService(repo)

			// Action
			err := sut.DeleteTransaction(context.Background(), tc.id)

			// Assert
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
