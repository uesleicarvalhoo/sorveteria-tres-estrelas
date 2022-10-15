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

	err = suite.db.AutoMigrate(repository.UserModel{})
	assert.NoError(suite.T(), err)
}

func (suite *PopsicleTestSuite) TeardownTest() {
	_ = suite.container.Terminate(suite.ctx)
}

func (suite *PopsicleTestSuite) TestCRUD() {
	repo := repository.NewUserPostgres(suite.db)

	storedUser := entity.User{
		ID:           entity.NewID(),
		Email:        "fakeuser@email.com",
		Name:         "Fake LastName",
		PasswordHash: "fakehash:123",
		Permissions:  []entity.Permission{entity.ReadWritePopsicle, entity.ReadWriteSalesRole},
	}

	suite.T().Run("test create a new user", func(t *testing.T) {
		err := repo.Create(suite.ctx, storedUser)
		assert.NoError(t, err)
	})

	suite.T().Run("test get user by id", func(t *testing.T) {
		found, err := repo.Get(suite.ctx, storedUser.ID)

		assert.NoError(t, err)

		assert.Equal(t, storedUser, found)
	})

	suite.T().Run("test get user by email", func(t *testing.T) {
		found, err := repo.GetByEmail(suite.ctx, storedUser.Email)

		assert.NoError(t, err)

		assert.Equal(t, storedUser, found)
	})

	suite.T().Run("test get all users", func(t *testing.T) {
		found, err := repo.GetAll(suite.ctx)

		assert.NoError(t, err)
		assert.Len(t, found, 1)
		assert.Equal(t, storedUser, found[0])
	})

	suite.T().Run("test Update should update sales_items table", func(t *testing.T) {
		// Arrange
		storedUser.Email = "new@email.com.br"

		// Action
		err := repo.Update(suite.ctx, storedUser)

		// Assert
		assert.NoError(t, err)

		found, err := repo.Get(suite.ctx, storedUser.ID)

		assert.NoError(t, err)
		assert.Equal(t, storedUser.Email, found.Email)
	})

	suite.T().Run("test delete user", func(t *testing.T) {
		err := repo.Delete(suite.ctx, storedUser.ID)

		assert.NoError(t, err)

		found, err := repo.Get(suite.ctx, storedUser.ID)

		assert.Equal(t, entity.User{}, found)
		assert.Error(t, err)
	})
}
