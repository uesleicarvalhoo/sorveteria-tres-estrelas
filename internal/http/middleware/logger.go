package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/logger"
)

func NewLogger(logger logger.Logger, serviceName, serviceVersion string) Middleware {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if strings.Contains(r.URL.Path, "/health") {
			next(w, r)

			return
		}

		next(w, r)

		statusCode := r.Response.StatusCode

		entry := map[string]interface{}{
			"log_version": "1.0.0",
			"date_time":   time.Now(),
			"product": map[string]interface{}{
				"name":        serviceName,
				"application": serviceName,
				"version":     serviceVersion,
				"http": map[string]string{
					"method": r.Method,
					"path":   r.URL.Path,
				},
			},
			"origin": map[string]interface{}{
				"application": serviceName,
				"ip":          r.RemoteAddr,
				"headers": map[string]string{
					"user_agent": r.UserAgent(),
					"origin":     r.Header.Get("Origin"),
					"refer":      r.Referer(),
				},
			},
			"context": map[string]interface{}{
				"service":     serviceName,
				"status_code": statusCode,
				"request_id":  r.Response.Header.Get("X-Request-Id"),
			},
		}

		switch {
		case statusCode >= http.StatusInternalServerError:
			logger.ErrorJSON(entry)
		case statusCode >= http.StatusBadRequest:
			logger.WarningJSON(entry)
		default:
			logger.InfoJSON(entry)
		}
	}
}
