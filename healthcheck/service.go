package healthcheck

import (
	"context"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/config"
)

type Service struct {
	cfg   *config.Config
	db    DatabasePing
	cache CachePing
}

var _ UseCase = Service{}

func NewService(cfg *config.Config, db DatabasePing, cache CachePing) Service {
	return Service{
		cfg:   cfg,
		db:    db,
		cache: cache,
	}
}

func (s Service) HealthCheck(ctx context.Context) HealthStatus {
	status := HealthStatus{
		Version: s.cfg.ServiceVersion,
		Status:  StatusUp,
	}

	if err := s.db.Ping(); err != nil {
		status.Database = StatusDown
		status.Status = StatusDown
	} else {
		status.Database = StatusUp
	}

	if err := s.cache.Ping(ctx); err != nil {
		status.Cache = StatusDown
		status.Status = StatusDown
	} else {
		status.Cache = StatusUp
	}

	return status
}
