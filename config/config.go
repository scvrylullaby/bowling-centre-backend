package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/scvrylullaby/bowling-centre-backend/pkg/logger"
)

type Config struct {
	APP struct {
		Env string
	}

	HTTP struct {
		Port string
		Host string
	}

	HTTP_CORS struct {
		Cors string
	}
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Log(".env file not found")
	}
	var cfg Config
	loadFromEnv(&cfg)
	return &cfg
}

func loadFromEnv(cfg *Config) {
	get := func(key string) string {
		return strings.TrimSpace(os.Getenv(key))
	}

	cfg.APP.Env = get("APP_ENV")

	cfg.HTTP.Host = get("HTTP_HOST")
	cfg.HTTP.Port = get("HTTP_PORT")

	cfg.HTTP_CORS.Cors = get("HTTP_CORS")
}
