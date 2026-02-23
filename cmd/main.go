package main

import (
	"net/http"

	"github.com/scvrylullaby/bowling-centre-backend/internal/config"
	"github.com/scvrylullaby/bowling-centre-backend/internal/middleware"
	"github.com/scvrylullaby/bowling-centre-backend/pkg/logger"
)

func main() {
	cfg := config.Load()

	logger.Init()
	
	mux := http.NewServeMux()
	handler := middleware.SetCors(cfg)(mux)

	logger.Log("Server has been started at %s:%s",cfg.HTTP.Host,cfg.HTTP.Port)
	http.ListenAndServe(":"+cfg.HTTP.Port, handler)
}