package healthcheck

import "context"

type UseCase interface {
	HealthCheck(ctx context.Context) HealthStatus
}

type DatabasePing interface {
	Ping() error
}

type CachePing interface {
	Ping(ctx context.Context) error
}
