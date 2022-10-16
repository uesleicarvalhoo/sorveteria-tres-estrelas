package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewLogrus(logger *logrus.Logger, serviceName, serviceVersion string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if strings.Contains(c.Path(), "/health") {
			return c.Next()
		}

		if err := c.Next(); err != nil {
			return err
		}

		statusCode := c.Response().StatusCode()
		reqHeaders := c.GetReqHeaders()
		respHeaders := c.GetRespHeaders()
		entry := logger.WithFields(logrus.Fields{
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
				"ip":          c.IP(),
				"headers": map[string]string{
					"user_agent": string(c.Context().UserAgent()),
					"origin":     reqHeaders["Origin"],
					"refer":      string(c.Context().Referer()),
				},
			},
			"context": map[string]interface{}{
				"service":     serviceName,
				"status_code": statusCode,
				"request_id":  respHeaders["X-Request-ID"],
			},
		})

		switch {
		case statusCode >= http.StatusInternalServerError:
			entry.Error()
		case statusCode >= http.StatusBadRequest:
			entry.Warning()
		default:
			entry.Info()
		}

		return nil
	}
}
