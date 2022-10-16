// @title Sorveteria três estrelas - API
// @version 1.0
// @description API para o cadastro de produtos, controle de vendas e fluxo de caixa para a sorveteria três estrelas

package main

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
	"github.com/sirupsen/logrus"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/api/handler"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/api/middleware"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/cache"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/config"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/database"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/logger"
)

const TIMEOUT = time.Second * 30

func gracefullShutdown(app *fiber.App) error {
	shutdownCh := make(chan error, 1)

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	go func() { shutdownCh <- app.Shutdown() }()

	select {
	case <-ctx.Done():
		return nil
	case err := <-shutdownCh:
		return err
	}
}

func NewFiber(appName, appVersion string, services *Services, logger *logrus.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               appName,
		DisableStartupMessage: true,
		ReadTimeout:           TIMEOUT,
		WriteTimeout:          TIMEOUT,
	})

	logrusMiddleware := middleware.NewLogrus(logger, appName, appVersion)

	app.Use(
		recover.New(),
		cors.New(),
		requestid.New(),
		logrusMiddleware,
	)

	handler.MakeHealthCheckRoutes(app)
	handler.MakeSwaggerRoutes(app.Group("/docs"))
	handler.MakeAuhtRoutes(app.Group("/auth"), services.authSvc)
	handler.MakeUserRoutes(app.Group("/users", middleware.NewAuth(services.authSvc, "users")), services.userSvc)
	handler.MakeSalesRoutes(app.Group("/sales", middleware.NewAuth(services.authSvc, "sales")), services.salesSvc)
	handler.MakePopsicleRoutes(
		app.Group("/popsicles", middleware.NewAuth(services.authSvc, "popsicles")), services.popsicleSvc)

	return app
}

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

	services := createServices(db, cache, cfg.SecretKey)

	// Http server
	app := NewFiber(cfg.ServiceName, cfg.ServiceVersion, services, logger)

	go func() { logger.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.HTTPPort))) }()

	logger.Infof("app started, listening on port %d", cfg.HTTPPort)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM)

	<-quit

	if err := gracefullShutdown(app); err != nil {
		logger.Fatalf("forcing app shutdown: %s", err)
	}
}
