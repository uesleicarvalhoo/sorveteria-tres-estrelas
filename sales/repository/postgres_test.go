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
	container *repository.PostgresContainer
	db        *gorm.DB
}

func TestPopsicleTestSuit(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(PopsicleTestSuite))
}

func (suite *PopsicleTestSuite) SetupTest() {
	var err error

	suite.ctx = context.Background()

	suite.container, err = repository.SetupPostgres(suite.ctx)
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

	s := sales.Sale{
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
		Date:        time.Now(),
	}

	suite.T().Run("test create sale", func(t *testing.T) {
		err := repo.Create(suite.ctx, s)

		assert.NoError(t, err)
	})

	suite.T().Run("test get sale", func(t *testing.T) {
		found, err := repo.Get(suite.ctx, s.ID)

		assert.NoError(t, err)

		assert.Equal(t, s.ID, found.ID)
		assert.Equal(t, s.Total, found.Total)
		assert.Equal(t, s.PaymentType, found.PaymentType)
		assert.Equal(t, s.Description, found.Description)
		assert.Len(t, found.Items, len(s.Items))
	})

	suite.T().Run("test get all sales", func(t *testing.T) {
		found, err := repo.GetAll(suite.ctx)

		assert.NoError(t, err)
		assert.Len(t, found, 1)
		assert.Equal(t, s.ID, found[0].ID)
		assert.Equal(t, s.Total, found[0].Total)
		assert.Equal(t, s.PaymentType, found[0].PaymentType)
		assert.Equal(t, s.Description, found[0].Description)
		assert.Equal(t, s.ID, found[0].ID)
		assert.Len(t, found[0].Items, len(s.Items))
	})
}
