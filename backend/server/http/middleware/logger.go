package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/logger"
)

func Logger(serviceName, serviceVersion string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if strings.Contains(c.Path(), "/health") {
			return c.Next()
		}

		if err := c.Next(); err != nil {
			return err
		}

		statusCode := c.Response().StatusCode()

		entry := map[string]interface{}{
			"log_version": "1.0.0",
			"date_time":   time.Now(),
			"product": map[string]interface{}{
				"name":        serviceName,
				"application": serviceName,
				"version":     serviceVersion,
				"http": map[string]string{
					"method": c.Method(),
					"path":   c.Path(),
				},
			},
			"origin": map[string]interface{}{
				"application": serviceName,
				"ip":          c.Context().RemoteAddr(),
				"headers": map[string]string{
					"user_agent": string(c.Context().UserAgent()),
					"origin":     c.GetRespHeader("Origin"),
					"refer":      string(c.Context().Referer()),
				},
			},
			"context": map[string]interface{}{
				"service":     serviceName,
				"status_code": statusCode,
				"request_id":  c.GetRespHeader("X-Request-Id"),
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

		return nil
	}
}
