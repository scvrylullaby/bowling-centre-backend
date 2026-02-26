package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/scvrylullaby/bowling-centre-backend/config"
	"github.com/scvrylullaby/bowling-centre-backend/internal/core"
	"github.com/scvrylullaby/bowling-centre-backend/internal/handlers"
	"github.com/scvrylullaby/bowling-centre-backend/internal/models"
	"github.com/scvrylullaby/bowling-centre-backend/pkg/logger"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cfg := config.Load()
	logger.Init()

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.HTTP_CORS.Cors},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	stateChan := make(chan models.DashboardState, 10)
	manager := core.NewManager(5, stateChan)

	go manager.Run()

	router.GET("/ws", gin.WrapH(handlers.Scoreboard(stateChan)))
	router.POST("/client", handlers.AddCustomer(manager))

	logger.Log("Server has been started at %s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	router.Run(":" + cfg.HTTP.Port)
}
