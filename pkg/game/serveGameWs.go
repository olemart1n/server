package game

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/olemart1n/server/pkg/game/utils"
)



func (m *Manager) ServeGameWS(w http.ResponseWriter, r *http.Request) {
	conn, err := utils.WebsocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	username := r.URL.Query().Get("username")
	spectator, id := NewSpectator(conn, username)
	m.Lock()
	defer m.Unlock()
	m.Spectators[spectator] = true


	// SEND ID
	jsonData, err := utils.CreateJsonObject("id", id)
	if err != nil {
		print(err)
	}
	spectator.Connection.WriteMessage(websocket.TextMessage, jsonData)

	//BROADCAST SPECTATOR TO ALL OTHERS
	m.broadcastPlayerToAll("spectator_joins", spectator, true)


// SEND SPECTATOR LIST
	spectators := getSpectatorList(m)
	data, _ := utils.CreateJsonObject("connected_spectators", spectators)
	if err := spectator.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Println("Error sending connected spectators")
	}



	// Create a context for this client's lifecycle
	ctx, cancel := context.WithCancel(context.Background())
	// Start client-specific goroutines
	go func() {
		spectator.handleConnection(m, cancel)
	}()
	go spectator.handlePingPong(m, ctx, cancel)
	go spectator.sendMessages(ctx, cancel)
}



func (m *Manager) removeSpectator (c *Client) {
		m.Lock()
		defer m.Unlock()
		if _, ok := m.Spectators[c]; ok {
			delete(m.Spectators, c)
			delete(m.Players, c)
			close(c.Egress)
			m.broadcastPlayerToAll("spectator_leaves", c, false)
			c.Connection.Close()
		}
	}
