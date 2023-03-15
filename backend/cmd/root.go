package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/config"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/logger"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/trace"
)

var rootCmd = &cobra.Command{
	Use:   "backend",
	Short: "backend server",
	Long:  "backend api for sales and cashflow control",
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

	if err := httpCommand.Execute(); err != nil {
		logger.Fatalf("error while executing command %v", err)
	}
}
