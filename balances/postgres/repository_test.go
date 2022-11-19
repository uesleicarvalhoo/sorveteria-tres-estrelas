//go:build integration || all

package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/balances/postgres"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/database"
	"gorm.io/gorm"
)

type BalancesPostgresTestSuite struct {
	suite.Suite
	ctx       context.Context //nolint:containedctx
	container *PostgresContainer
	db        *gorm.DB
}

func TestBalancesPostgresTestSuit(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(BalancesPostgresTestSuite))
}

func (suite *BalancesPostgresTestSuite) SetupTest() {
	var err error

	suite.ctx = context.Background()

	suite.container, err = SetupPostgres(suite.ctx)
	assert.NoError(suite.T(), err)

	suite.db, err = database.NewPostgresConnection(
		suite.container.Username,
		suite.container.Password,
		suite.container.Database,
		suite.container.Host,
		suite.container.Port)
	assert.NoError(suite.T(), err)

	err = suite.db.AutoMigrate(&balances.Balance{})
	assert.NoError(suite.T(), err)
}

func (suite *BalancesPostgresTestSuite) TeardownTest() {
	_ = suite.container.Terminate(suite.ctx)
}

func (suite *BalancesPostgresTestSuite) TestCRUD() {
	repo := postgres.NewRepository(suite.db)

	balance := balances.Balance{
		ID:          uuid.New(),
		Description: "new sale",
		Operation:   balances.OperationSale,
		Value:       1.20,
		CreatedAt:   time.Now(),
	}

	suite.T().Run("test create balance", func(t *testing.T) {
		err := repo.Create(suite.ctx, balance)
		assert.NoError(t, err)
	})

	suite.T().Run("test get balance", func(t *testing.T) {
		found, err := repo.Get(suite.ctx, balance.ID)

		assert.NoError(t, err)
		assert.Equal(t, balance.ID, found.ID)
	})

	suite.T().Run("test get all balance", func(t *testing.T) {
		found, err := repo.GetAll(suite.ctx)

		assert.NoError(t, err)
		assert.Len(t, found, 1)
		assert.Equal(t, found[0].ID, balance.ID)
		assert.Equal(t, found[0].Description, balance.Description)
		assert.Equal(t, found[0].Operation, balance.Operation)
	})

	suite.T().Run("test update balance", func(t *testing.T) {
		balance.Description = "venda de uma caixa de picoles"

		err := repo.Update(suite.ctx, &balance)

		assert.NoError(t, err)

		updated, err := repo.Get(suite.ctx, balance.ID)
		assert.NoError(t, err)

		assert.Equal(t, balance.Description, updated.Description)
	})

	suite.T().Run("test delete balance", func(t *testing.T) {
		err := repo.Delete(suite.ctx, balance.ID)

		assert.NoError(t, err)
	})
}
