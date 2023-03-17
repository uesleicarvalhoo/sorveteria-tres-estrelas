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

			svc := healthcheck.NewService(cfg, dbMock)

			// Action
			status := svc.HealthCheck(context.Background())

			// Assert
			assert.Equal(t, tc.expectedStatus, status)
		})
	}
}
