//go:build integration || all

package repository_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/database"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/popsicle"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/popsicle/repository"
	"gorm.io/gorm"
)

type PopsicleTestSuite struct {
	suite.Suite
	ctx       context.Context //nolint:containedctx
	container *PostgresContainer
	db        *gorm.DB
}

func TestPopsicleTestSuit(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(PopsicleTestSuite))
}

func (suite *PopsicleTestSuite) SetupTest() {
	var err error

	suite.ctx = context.Background()

	suite.container, err = SetupPostgres(suite.ctx)
	assert.NoError(suite.T(), err)

	suite.db, err = database.NewPostgreConnection(
		suite.container.Username,
		suite.container.Password,
		suite.container.Database,
		suite.container.Host,
		suite.container.Port)
	assert.NoError(suite.T(), err)

	err = suite.db.AutoMigrate(&popsicle.Popsicle{})
	assert.NoError(suite.T(), err)
}

func (suite *PopsicleTestSuite) TeardownTest() {
	_ = suite.container.Terminate(suite.ctx)
}

func (suite *PopsicleTestSuite) TestCRUD() {
	repo := repository.NewPostgresRepository(suite.db)

	pop := popsicle.Popsicle{
		ID:     uuid.New(),
		Flavor: "morango",
		Price:  1.20,
	}

	suite.T().Run("test create popsicle", func(t *testing.T) {
		err := repo.Create(suite.ctx, pop)
		assert.NoError(t, err)
	})

	suite.T().Run("test get popsicle", func(t *testing.T) {
		found, err := repo.Get(suite.ctx, pop.ID)

		assert.NoError(t, err)
		assert.Equal(t, pop, found)
	})

	suite.T().Run("test get all popsicle", func(t *testing.T) {
		found, err := repo.GetAll(suite.ctx)

		assert.NoError(t, err)
		assert.Len(t, found, 1)
		assert.Equal(t, found[0], pop)
	})

	suite.T().Run("test update popsicle", func(t *testing.T) {
		pop.Flavor = "coco com morango"

		err := repo.Update(suite.ctx, &pop)

		assert.NoError(t, err)
	})

	suite.T().Run("test delete popsicle", func(t *testing.T) {
		err := repo.Delete(suite.ctx, pop.ID)

		assert.NoError(t, err)
	})
}
