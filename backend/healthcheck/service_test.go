package healthcheck_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/config"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/healthcheck"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/healthcheck/mocks"
)

func TestService(t *testing.T) {
	t.Parallel()

	tests := []struct {
		about          string
		cacheErr       error
		dbErr          error
		version        string
		expectedStatus healthcheck.HealthStatus
	}{
		{
			about:   "when all services are ok",
			version: "0.0.0",
			expectedStatus: healthcheck.HealthStatus{
				Version:  "0.0.0",
				Status:   healthcheck.StatusUp,
				Database: healthcheck.StatusUp,
				Cache:    healthcheck.StatusUp,
			},
		},
		{
			about:   "when database is down",
			dbErr:   errors.New("db error"),
			version: "0.0.0",
			expectedStatus: healthcheck.HealthStatus{
				Version:  "0.0.0",
				Status:   healthcheck.StatusDown,
				Database: healthcheck.StatusDown,
				Cache:    healthcheck.StatusUp,
			},
		},
		{
			about:    "when cache is down",
			cacheErr: errors.New("cache error"),
			version:  "0.0.0",
			expectedStatus: healthcheck.HealthStatus{
				Version:  "0.0.0",
				Status:   healthcheck.StatusDown,
				Database: healthcheck.StatusUp,
				Cache:    healthcheck.StatusDown,
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.about, func(t *testing.T) {
			t.Parallel()

			// Arrange
			dbMock := mocks.NewDatabasePing(t)
			dbMock.On("Ping").Return(tc.dbErr)

			cfg := &config.Config{
				ServiceVersion: tc.version,
			}

			cacheMock := mocks.NewCachePing(t)
			cacheMock.On("Ping", context.Background()).Return(tc.cacheErr)
			svc := healthcheck.NewService(cfg, dbMock, cacheMock)

			// Action
			status := svc.HealthCheck(context.Background())

			// Assert
			assert.Equal(t, tc.expectedStatus, status)
		})
	}
}
