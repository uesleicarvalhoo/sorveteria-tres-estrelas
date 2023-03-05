// @title Sorveteria três estrelas - Backend API
// @version 1.0
// @description API para o cadastro de produtos, controle de vendas e fluxo de caixa para a sorveteria três estrelas

package main

import (
	"context"

	"github.com/kong/go-kong/kong"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/cmd/api/fiber"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/config"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/infrastructure/cache"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/infrastructure/database"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/infrastructure/http"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/ioc"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/logger"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/trace"
)

func main() {
	cfg, err := config.NewFromEnv()
	if err != nil {
		panic(err)
	}

	logger, err := logger.NewLogrus(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	// Dependencies
	db, err := database.NewPostgresConnectionWithMigration(
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)
	if err != nil {
		logger.Fatalf("couldn't connect to database: %s", err)
	}

	cache, err := cache.NewRedis(cfg.CacheURI, cfg.CachePassword)
	if err != nil {
		logger.Fatalf("couldn't connect to redis: %s", err)
	}

	provider, err := trace.NewProvider(
		trace.ProviderConfig{
			Endpoint:       cfg.TraceEndpoint,
			ServiceName:    cfg.ServiceName,
			ServiceVersion: cfg.ServiceVersion,
			Environment:    cfg.Environment,
			Disabled:       cfg.TraceEnabled,
		})
	if err != nil {
		logger.Fatalf("couldn't connect to provider: %s", err)
	}
	defer provider.Close(context.Background())

	kong, err := kong.NewClient(&cfg.KongURL, nil)
	if err != nil {
		logger.Fatalf("couldn't connect to kong: %s", err)
	}

	con, _ := db.DB()

	healthSvc := ioc.NewHealthCheckService(cfg, con, cache)
	saleSvc := ioc.NewSaleService(db)
	productSvc := ioc.NewProductService(db)
	usersSvc := ioc.NewUserService(db)
	paymentSvc := ioc.NewPaymentService(db)
	cashflowSvc := ioc.NewCashFlowService(db)
	authSvc := ioc.NewAuthService(db, cache, kong, cfg.SecretKey, cfg.KongConsumer, cfg.KongJwtKey)

	h := fiber.Handlers(
		cfg.ServiceName,
		cfg.ServiceVersion,
		logger,
		authSvc,
		healthSvc,
		usersSvc,
		productSvc,
		saleSvc,
		paymentSvc,
		cashflowSvc,
	)

	if err := http.Start(cfg.HTTPPort, cfg.ServiceName, cfg.ServiceVersion, h, logger); err != nil {
		panic(err)
	}
}