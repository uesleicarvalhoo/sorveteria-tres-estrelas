// @title Sorveteria três estrelas - API
// @version 1.0
// @description API para o cadastro de produtos, controle de vendas e fluxo de caixa para a sorveteria três estrelas

package main

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cache"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/config"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/database"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/http"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/http/fiber"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/ioc"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/logger"
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

	authSvc := ioc.NewAuthService(cfg.SecretKey, db, cache)
	saleSvc := ioc.NewSaleService(db)
	productSvc := ioc.NewProductService(db)
	usersSvc := ioc.NewUserService(db)
	paymentSvc := ioc.NewPaymentService(db)
	cashflowSvc := ioc.NewCashFlowService(db)

	h := fiber.Handlers(
		cfg.ServiceName, cfg.ServiceVersion, authSvc, usersSvc, productSvc, saleSvc, paymentSvc, cashflowSvc)

	if err := http.Start(cfg.HTTPPort, cfg.ServiceName, cfg.ServiceVersion, h, logger); err != nil {
		panic(err)
	}
}
