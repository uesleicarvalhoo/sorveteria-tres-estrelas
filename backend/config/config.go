package config

import "github.com/netflix/go-env"

type Config struct {
	// General
	ServiceName    string `env:"SERVICE_NAME,default=sorveteria-tres-estrelas"`
	ServiceVersion string `env:"SERVICE_VERSION,default=0.0.0"`
	SecretKey      string `env:"SECRET_KEY,default=my-secret-key"`
	LogLevel       string `env:"LOG_LEVEL,default=INFO"`
	Environment    string `env:"ENVIRONMENT,default=dev"`

	// HTTP Server
	HTTPPort int `env:"HTTP_SERVER_PORT,default=5000"`

	// Database
	DBHost     string `env:"DATABASE_HOST,default=localhost"`
	DBPort     int    `env:"DATABASE_PORT,default=5432"`
	DBName     string `env:"DATABASE_NAME,default=sorveteria-tres-estrelas"`
	DBUser     string `env:"DATABASE_USER,default=postgres"`
	DBPassword string `env:"DATABASE_PASSWORD,default=secret"`

	// Tracer
	TraceEndpoint string `env:"TRACER_ENDPOINT,default=http://localhost:9411/api/v2/spans"`
	TraceEnabled  bool   `env:"TRACE_ENABLED,default=true"`

	// Kong
	KongURL      string `env:"KONG_URL,default=http://localhost:8001"`
	KongConsumer string `env:"KONG_CONSUMER,default=auth-consumer"`
	KongJwtKey   string `env:"KONG_JWT_KEY,default=auth-jwt"`
}

func NewFromEnv() (*Config, error) {
	cfg := &Config{}

	if _, err := env.UnmarshalFromEnviron(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
