package ioc

import "github.com/uesleicarvalhoo/sorveteria-tres-estrelas/healthcheck"

func NewHealthCheckService(db healthcheck.DatabasePing, cache healthcheck.CachePing) healthcheck.Service {
	return healthcheck.NewService(db, cache)
}
