//go:build integration || all

package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/database"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/sales/repository"
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

	err = suite.db.AutoMigrate(&repository.SaleModel{}, &repository.SaleItemModel{})
	assert.NoError(suite.T(), err)
}

func (suite *PopsicleTestSuite) TeardownTest() {
	_ = suite.container.Terminate(suite.ctx)
}

func (suite *PopsicleTestSuite) TestCRUD() {
	repo := repository.NewPostgresRepository(suite.db)

	sale := sales.Sale{
		ID:          uuid.New(),
		PaymentType: sales.AnotherPayments,
		Total:       8.75,
		Items: []sales.SaleItem{
			{
				Name:   "Picole de coco",
				Amount: 10,
				Price:  0.5,
			},
			{
				Name:   "Picole de chocolate",
				Amount: 5,
				Price:  0.75,
			},
		},
		Description: "i'm a sale description",
		Date:        time.Date(2022, 10, 13, 0, 0, 0, 0, time.Local),
	}

	suite.T().Run("test create sale", func(t *testing.T) {
		err := repo.Create(suite.ctx, sale)
		assert.NoError(t, err)
	})

	suite.T().Run("test get sale", func(t *testing.T) {
		found, err := repo.Get(suite.ctx, sale.ID)

		assert.NoError(t, err)

		assert.Equal(t, sale.ID, found.ID)
		assert.Equal(t, sale.Total, found.Total)
		assert.Equal(t, sale.PaymentType, found.PaymentType)
		assert.Equal(t, sale.Description, found.Description)
		assert.Len(t, found.Items, len(sale.Items))
	})

	suite.T().Run("test get all sales", func(t *testing.T) {
		found, err := repo.GetAll(suite.ctx)

		assert.NoError(t, err)
		assert.Len(t, found, 1)
		assert.Equal(t, sale.ID, found[0].ID)
		assert.Equal(t, sale.Total, found[0].Total)
		assert.Equal(t, sale.PaymentType, found[0].PaymentType)
		assert.Equal(t, sale.Description, found[0].Description)
		assert.Equal(t, sale.ID, found[0].ID)
		assert.Len(t, found[0].Items, len(sale.Items))
	})

	suite.T().Run("test Update should update sales_items table", func(t *testing.T) {
		// Arrange
		sale.Items = []sales.SaleItem{
			{Name: "Picole de amendoin", Amount: 1, Price: 1.0},
		}

		// Action
		err := repo.Update(suite.ctx, sale)

		// Assert
		assert.NoError(t, err)

		found, err := repo.Get(suite.ctx, sale.ID)

		assert.NoError(t, err)
		assert.Equal(t, sale.Items, found.Items)
	})

	suite.T().Run("test delete sale", func(t *testing.T) {
		err := repo.Delete(suite.ctx, sale.ID)

		assert.NoError(t, err)

		found, err := repo.Get(suite.ctx, sale.ID)

		assert.Equal(t, sales.Sale{}, found)
		assert.Error(t, err)
	})

	suite.T().Run("test search should return only dates between start and end date", func(t *testing.T) {
		// Arrange
		storedSales := []sales.Sale{
			{
				ID:          uuid.New(),
				PaymentType: sales.AnotherPayments,
				Total:       5,
				Items: []sales.SaleItem{
					{Name: "Picole de coco", Amount: 10, Price: 0.5},
				},
				Description: "i'm a sale description",
				Date:        time.Date(2022, 10, 13, 0, 0, 0, 0, time.Local),
			},
			{
				ID:          uuid.New(),
				PaymentType: sales.AnotherPayments,
				Total:       3.75,
				Items: []sales.SaleItem{
					{Name: "Picole de chocolate", Amount: 5, Price: 0.75},
				},
				Description: "i'm a sale description",
				Date:        time.Date(2022, 10, 11, 0, 0, 0, 0, time.Local),
			},
			{
				ID:          uuid.New(),
				PaymentType: sales.AnotherPayments,
				Total:       5,
				Items: []sales.SaleItem{
					{Name: "Picole de morango", Amount: 5, Price: 1},
				},
				Description: "i'm a sale description",
				Date:        time.Date(2022, 10, 31, 0, 0, 0, 0, time.Local),
			},
			{
				ID:          uuid.New(),
				PaymentType: sales.AnotherPayments,
				Total:       6.25,
				Items: []sales.SaleItem{
					{Name: "Picole de coco com goiaba", Amount: 5, Price: 1.25},
				},
				Description: "i'm a sale description",
				Date:        time.Date(2022, 9, 1, 0, 0, 0, 0, time.Local),
			},
		}

		for _, s := range storedSales {
			err := repo.Create(suite.ctx, s)

			assert.NoError(t, err)
		}

		start := time.Date(2022, 10, 1, 0, 0, 0, 0, time.Local)
		end := time.Date(2022, 10, 31, 23, 59, 59, 0, time.Local)

		// Action
		found, err := repo.Search(suite.ctx, start, end)

		// Assert
		assert.NoError(t, err)
		assert.Len(t, found, 3)
	})
}
