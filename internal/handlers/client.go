package handlers

import (
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/scvrylullaby/bowling-centre-backend/internal/core"
	"github.com/scvrylullaby/bowling-centre-backend/pkg/logger"
)

var сounter int64

type ClientRequest struct {
	PlayTime int `json:"play_time"`
	Timeout  int `json:"timeout"`
}

func AddCustomer(manager *core.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ClientRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			logger.Log("Invalid JSON: ", err)
			return
		}

		newID := atomic.AddInt64(&сounter, 1)
		pTime := time.Duration(req.PlayTime) * time.Second
		tOut := time.Duration(req.Timeout) * time.Second
		go core.GenerateClient(
			int(newID),
			manager,
			pTime,
			tOut,
		)

		c.JSON(http.StatusAccepted, gin.H{
			"status":      "customer_queued",
			"customer_id": newID,
		})
	}
}
