package cmd

import (
	"context"

	"github.com/kong/go-kong/kong"
	"github.com/spf13/cobra"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/config"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/database"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/logger"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/trace"
	"gorm.io/gorm"
)

type application struct {
	db     *gorm.DB
	config *config.Config
	kong   *kong.Client
}

func (app *application) execute() error {
	cmd := &cobra.Command{
		Use:   "backend",
		Short: "backend server",
		Long:  "backend api for sales and cashflow control",
	}

	cmd.AddCommand(app.httpServer())

	return cmd.Execute()
}

func Execute() {
	defer func() {
		if err := recover(); err != nil {
			logger.Fatalf("unexpected error while executing command %v", err)
		}
	}()

	cfg, err := config.NewFromEnv()
	if err != nil {
		panic(err)
	}

	if err := logger.Configure(cfg.LogLevel); err != nil {
		panic(err)
	}

	provider, err := trace.NewProvider(
		trace.ProviderConfig{
			Endpoint:       cfg.TraceEndpoint,
			ServiceName:    cfg.ServiceName,
			ServiceVersion: cfg.ServiceVersion,
			Environment:    cfg.Environment,
			Disabled:       !cfg.TraceEnabled,
		})
	if err != nil {
		logger.Fatalf("couldn't connect to provider: %s", err)
	}
	defer provider.Close(context.Background())

	// Dependencies
	db, err := database.NewPostgresConnectionWithMigration(cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)
	if err != nil {
		logger.Fatalf("error when connect to database: %s", err)
	}

	kong, err := kong.NewClient(&cfg.KongURL, nil)
	if err != nil {
		logger.Fatalf("error when connect to kong: %s", err)
	}

	app := application{
		db:     db,
		kong:   kong,
		config: cfg,
	}

	if err := app.execute(); err != nil {
		logger.Fatalf("error while executing command: %v", err)
	}
}
