package core

import (
	"time"

	"github.com/scvrylullaby/bowling-centre-backend/internal/models"
)

func RunLane(id int, input chan *models.Client, done chan int) {
	for customer := range input {
		time.Sleep(customer.PlayTime)
		done <- id
	}
}
