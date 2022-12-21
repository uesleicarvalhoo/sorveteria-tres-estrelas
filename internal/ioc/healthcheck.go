package ioc

import (
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/config"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/healthcheck"
)

func NewHealthCheckService(
	cfg *config.Config, db healthcheck.DatabasePing, cache healthcheck.CachePing,
) healthcheck.Service {
	return healthcheck.NewService(cfg, db, cache)
}
