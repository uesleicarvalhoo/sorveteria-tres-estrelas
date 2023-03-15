package healthcheck

import (
	"context"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/config"
)

type Service struct {
	cfg *config.Config
	db  DatabasePing
}

var _ UseCase = Service{}

func NewService(cfg *config.Config, db DatabasePing) Service {
	return Service{
		cfg: cfg,
		db:  db,
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

	return status
}
