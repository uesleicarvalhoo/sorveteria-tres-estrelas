// @title Sorveteria três estrelas - API
// @version 1.0
// @description API para o cadastro de produtos, controle de vendas e fluxo de caixa para a sorveteria três estrelas

package main

import (
	"context"
	"log"
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
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/database"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/logger"
)

const TIMEOUT = time.Second * 30

type Options func(app *fiber.App) error

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

func NewFiber(services *Services, logger *logrus.Logger, options ...Options) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		AppName:               "sorveteria-tres-estrelas",
		DisableStartupMessage: true,
		ReadTimeout:           TIMEOUT,
		WriteTimeout:          TIMEOUT,
	})

	authMiddleware := middleware.NewAuth(services.authSvc)
	logrusMiddleware := middleware.NewLogrus(logger, "sorveteria-tres-estrelas", "0.0.0")

	app.Use(
		recover.New(),
		cors.New(),
		requestid.New(),
		logrusMiddleware,
	)

	handler.MakePopsicleRoutes(app.Group("/popsicles", authMiddleware), services.popsicleSvc)
	handler.MakeSalesRoutes(app.Group("/sales", authMiddleware), services.salesSvc)
	handler.MakeUserRoutes(app.Group("/users", authMiddleware), services.userSvc)
	handler.MakeAuhtRoutes(app.Group("/auth"), services.authSvc)
	handler.MakeHealthCheckRoutes(app)
	handler.MakeSwaggerRoutes(app)

	for _, op := range options {
		if err := op(app); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func main() {
	logger, err := logger.NewLogrus("debug")
	if err != nil {
		panic(err)
	}

	dbPort := 5432

	db, err := database.NewPostgresConnectionWithMigration(
		"postgres", "secret", "sorveteria-tres-estrelas", "localhost", dbPort)
	if err != nil {
		panic(err)
	}

	cache, err := cache.NewRedis("localhost:6379", "")
	if err != nil {
		panic(err)
	}

	services := createServices(db, cache, "my-secret-key")

	app, err := NewFiber(services, logger)
	if err != nil {
		panic(err)
	}

	go log.Fatal(app.Listen(":8080"))

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM)

	<-quit

	if err := gracefullShutdown(app); err != nil {
		panic(err)
	}
}
