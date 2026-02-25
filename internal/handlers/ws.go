package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/scvrylullaby/bowling-centre-backend/internal/models"
	"github.com/scvrylullaby/bowling-centre-backend/pkg/logger"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Scoreboard(stateChan chan models.DashboardState) http.HandlerFunc {
	clients := make(map[*websocket.Conn]bool)
	go func() {
		for state := range stateChan {
			for client := range clients {
				err := client.WriteJSON(state)
				if err != nil {
					client.Close()
					delete(clients, client)
				}
			}
		}
	}()
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Log("Websocket upgrade error: ", err)
			return
		}
		clients[ws] = true
	}
}
