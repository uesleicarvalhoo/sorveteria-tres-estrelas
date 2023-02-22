//go:build integration || all

package postgres_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/database"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/product"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/product/postgres"
	"gorm.io/gorm"
)

type ProductsPostgresTestSuite struct {
	suite.Suite
	ctx       context.Context //nolint:containedctx
	container *PostgresContainer
	db        *gorm.DB
}

func TestProductsPostgresTestSuit(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ProductsPostgresTestSuite))
}

func (suite *ProductsPostgresTestSuite) SetupTest() {
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

	err = suite.db.AutoMigrate(&product.Product{})
	assert.NoError(suite.T(), err)
}

func (suite *ProductsPostgresTestSuite) TeardownTest() {
	_ = suite.container.Terminate(suite.ctx)
}

func (suite *ProductsPostgresTestSuite) TestCRUD() {
	repo := postgres.NewRepository(suite.db)

	product := product.Product{
		ID:            uuid.New(),
		Name:          "picole de morango",
		PriceVarejo:   1.20,
		PriceAtacado:  1,
		AtacadoAmount: 10,
	}

	suite.T().Run("test create product", func(t *testing.T) {
		err := repo.Create(suite.ctx, product)
		assert.NoError(t, err)
	})

	suite.T().Run("test get product", func(t *testing.T) {
		found, err := repo.Get(suite.ctx, product.ID)

		assert.NoError(t, err)
		assert.Equal(t, product, found)
	})

	suite.T().Run("test get all product", func(t *testing.T) {
		found, err := repo.GetAll(suite.ctx)

		assert.NoError(t, err)
		assert.Len(t, found, 1)
		assert.Equal(t, found[0], product)
	})

	suite.T().Run("test update product", func(t *testing.T) {
		product.Name = "picole de coco com morango"

		err := repo.Update(suite.ctx, &product)

		assert.NoError(t, err)

		updated, err := repo.Get(suite.ctx, product.ID)
		assert.NoError(t, err)

		assert.Equal(t, product.Name, updated.Name)
	})

	suite.T().Run("test delete product", func(t *testing.T) {
		err := repo.Delete(suite.ctx, product.ID)

		assert.NoError(t, err)
	})
}
