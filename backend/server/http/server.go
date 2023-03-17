package http

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/kong/go-kong/kong"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/config"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/ioc"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/logger"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/server/http/middleware"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/server/http/routes"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/shutdown"
	"gorm.io/gorm"
)

// @title Sorveteria três estrelas - Backend API
// @version 1.0
// @description API para o cadastro de produtos, controle de vendas e fluxo de caixa para a sorveteria três estrelas
func StartServer(cfg *config.Config, db *gorm.DB, kong *kong.Client) {
	con, err := db.DB()
	if err != nil {
		logger.Fatalf("error when getting db connection: %s", err)
	}

	// Services
	healthSvc := ioc.NewHealthCheckService(cfg, con)
	salesSvc := ioc.NewSaleService(db)
	productSvc := ioc.NewProductService(db)
	userSvc := ioc.NewUserService(db)
	paymentSvc := ioc.NewPaymentService(db)
	cashflowSvc := ioc.NewCashFlowService(db)
	authSvc := ioc.NewAuthService(db, kong, cfg.SecretKey, cfg.KongConsumer, cfg.KongJwtKey)

	// Server
	app := fiber.New(fiber.Config{
		AppName:               cfg.ServiceName,
		DisableStartupMessage: true,
	})

	app.Use(
		recover.New(),
		cors.New(),
		requestid.New(),
		middleware.Otel(cfg.ServiceName),
		middleware.Logger(cfg.ServiceName, cfg.ServiceVersion),
	)

	routes.HealthCheck(app, healthSvc)
	routes.Swagger(app.Group("/docs"))
	routes.Auth(app.Group("/auth"), authSvc)
	routes.User(app.Group("/user"), userSvc)
	routes.Sales(app.Group("/sales"), salesSvc)
	routes.Products(app.Group("/products"), productSvc)
	routes.Payments(app.Group("/payments"), paymentSvc)
	routes.CashFlow(app.Group("/cashflow"), cashflowSvc)

	// Graceful shutdown
	shutdown.Subscribe(func(ctx context.Context) error {
		logger.Infof("http server running on port: %d", cfg.HTTPPort)

		return app.Listen(fmt.Sprintf(":%d", cfg.HTTPPort))
	}, func(ctx context.Context) error {
		if err := app.Server().ShutdownWithContext(ctx); err != nil {
			logger.Errorf("graceful shutdown failed: %s", err)

			return err
		}

		return nil
	})
}
