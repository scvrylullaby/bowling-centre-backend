package core

import (
	"time"

	"github.com/scvrylullaby/bowling-centre-backend/internal/models"
)

func GenerateClient(id int, m *Manager, playTime, timeout time.Duration) {
	c := &models.Client{
		ID:       id,
		PlayTime: playTime,
		Timeout:  timeout,
		Start:    make(chan struct{}),
	}

	m.Incoming <- c
	select {
	case <-c.Start:
	case <-time.After(c.Timeout):
		m.LeftQueue <- c.ID
	}
}
