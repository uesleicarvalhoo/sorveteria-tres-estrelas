package config

import "github.com/netflix/go-env"

type Config struct {
	// General
	ServiceName    string `env:"SERVICE_NAME,default=sorveteria-tres-estrelas"`
	ServiceVersion string `env:"SERVICE_VERSION,default=0.0.0"`
	SecretKey      string `env:"SECRET_KEY,default=my-secret-key"`
	LogLevel       string `env:"LOG_LEVEL,default=INFO"`

	// HTTP Server
	HTTPPort int `env:"HTTP_SERVER_PORT,default=8080"`

	// Database
	DBHost     string `env:"DATABASE_HOST,default=localhost"`
	DBPort     int    `env:"DATABASE_PORT,default=5432"`
	DBName     string `env:"DATABASE_NAME,default=sorveteria-tres-estrelas"`
	DBUser     string `env:"DATABASE_USER,default=postgres"`
	DBPassword string `env:"DATABASE_PASSWORD,default=secret"`

	// Cache
	CacheURI      string `env:"CACHE_URI,default=localhost:6379"`
	CachePassword string `env:"CACHE_PASSWORD"`
}

func NewFromEnv() (*Config, error) {
	cfg := &Config{}

	if _, err := env.UnmarshalFromEnviron(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
