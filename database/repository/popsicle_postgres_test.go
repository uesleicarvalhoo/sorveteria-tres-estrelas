//go:build integration || all

package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/database"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/database/repository"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"gorm.io/gorm"
)

type PopsiclePostgresTestSuite struct {
	suite.Suite
	ctx       context.Context //nolint:containedctx
	container *PostgresContainer
	db        *gorm.DB
}

func TestPopsiclePostgresTestSuit(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(PopsiclePostgresTestSuite))
}

func (suite *PopsiclePostgresTestSuite) SetupTest() {
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

	err = suite.db.AutoMigrate(&repository.PopsicleModel{})
	assert.NoError(suite.T(), err)
}

func (suite *PopsiclePostgresTestSuite) TeardownTest() {
	_ = suite.container.Terminate(suite.ctx)
}

func (suite *PopsiclePostgresTestSuite) TestCRUD() {
	repo := repository.NewPopsiclePostgres(suite.db)

	pop := entity.Popsicle{
		ID:     entity.NewID(),
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
