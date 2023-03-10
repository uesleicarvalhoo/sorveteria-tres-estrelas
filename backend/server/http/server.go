package http

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/kong/go-kong/kong"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/config"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/infrastructure/cache"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/infrastructure/database"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/ioc"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/logger"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/server/http/middleware"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/server/http/routes"
)

// @title Sorveteria três estrelas - Backend API
// @version 1.0
// @description API para o cadastro de produtos, controle de vendas e fluxo de caixa para a sorveteria três estrelas
func StartServer() {
	// Dependencies
	cfg, err := config.NewFromEnv()
	if err != nil {
		logger.Fatalf("error when reading config: %s", err)
	}

	db, err := database.NewPostgresConnectionWithMigration(cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)
	if err != nil {
		logger.Fatalf("error when connect to database: %s", err)
	}

	cache, err := cache.NewRedis(cfg.CacheURI, cfg.CachePassword)
	if err != nil {
		logger.Fatalf("error when connect to redis: %s", err)
	}

	kong, err := kong.NewClient(&cfg.KongURL, nil)
	if err != nil {
		logger.Fatalf("error when connect to kong: %s", err)
	}

	con, err := db.DB()
	if err != nil {
		logger.Fatalf("error when getting db connection: %s", err)
	}

	// Services
	healthSvc := ioc.NewHealthCheckService(cfg, con, cache)
	salesSvc := ioc.NewSaleService(db)
	productSvc := ioc.NewProductService(db)
	userSvc := ioc.NewUserService(db)
	paymentSvc := ioc.NewPaymentService(db)
	cashflowSvc := ioc.NewCashFlowService(db)
	authSvc := ioc.NewAuthService(db, cache, kong, cfg.SecretKey, cfg.KongConsumer, cfg.KongJwtKey)

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

	go func() {
		logger.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.HTTPPort)))
	}()

	logger.Infof("http server running on port: %d", cfg.HTTPPort)

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGALRM)

	<-quit

	if err := app.Server().ShutdownWithContext(ctx); err != nil {
		logger.Errorf("graceful shutdown failed: %s", err)
	}
}
