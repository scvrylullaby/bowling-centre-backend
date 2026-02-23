package config

import (
	"net/http"

	"github.com/rs/cors"
)

func SetCors(cfg *Config) func(http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{cfg.HTTP_CORS.Cors},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler
}