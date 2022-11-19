// @title Sorveteria três estrelas - API
// @version 1.0
// @description API para o cadastro de produtos, controle de vendas e fluxo de caixa para a sorveteria três estrelas

package main

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cache"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/config"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/fiber"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/database"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/ioc"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/logger"
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
	balanceSvc := ioc.NewBalanceService(db)

	h := fiber.Handlers(
		cfg.ServiceName, cfg.ServiceVersion, authSvc, usersSvc, productSvc, saleSvc, balanceSvc)

	if err := api.Start(cfg.HTTPPort, cfg.ServiceName, cfg.ServiceVersion, h, logger); err != nil {
		panic(err)
	}
}
