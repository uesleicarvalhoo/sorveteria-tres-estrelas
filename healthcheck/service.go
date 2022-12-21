package healthcheck

import "context"

type Service struct {
	db    DatabasePing
	cache CachePing
}

var _ UseCase = Service{}

func NewService(db DatabasePing, cache CachePing) Service {
	return Service{
		db:    db,
		cache: cache,
	}
}

func (s Service) HealthCheck(ctx context.Context) HealthStatus {
	status := HealthStatus{
		App: StatusUp,
	}

	if err := s.db.Ping(); err != nil {
		status.Database = StatusDown
		status.App = StatusDown
	} else {
		status.Database = StatusUp
	}

	if err := s.cache.Ping(ctx); err != nil {
		status.Cache = StatusDown
		status.App = StatusDown
	} else {
		status.Cache = StatusUp
	}

	return status
}
