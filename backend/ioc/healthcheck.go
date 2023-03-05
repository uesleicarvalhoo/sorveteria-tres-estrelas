package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/config"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/healthcheck"
)

func NewHealthCheckService(
	cfg *config.Config, db healthcheck.DatabasePing, cache healthcheck.CachePing,
) healthcheck.Service {
	return healthcheck.NewService(cfg, db, cache)
}
