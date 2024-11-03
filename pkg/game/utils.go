package game

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)
func sendConnectedPlayersLength (c *GameClient) {
	data := map[string]interface{}{
		"name": "connectedPlayersLength",
		"payload": map[string]interface{}{
			"value": len(c.gameManager.gameClients),
		},
	}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}
	err = c.connection.WriteMessage(websocket.TextMessage, jsonBytes)
	if err != nil {
		log.Printf("Error sending client count to %s: %v", c.connection.RemoteAddr(), err)
	}
}

func (m *GameManager) broadcastPlayer(name string, c *GameClient) {
	data := map[string]interface{}{
		"name": name,
		"payload": map[string]interface{}{
			"value": c.username,
		},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}
	for player := range m.gameClients {
		if player != c {
			if err := player.connection.WriteMessage(websocket.TextMessage, jsonData); err != nil {
				log.Println(err)
				return
			}
		}
		
	}
}

func (m *GameManager) broadcastNewPlayer( c *GameClient) {
	m.broadcastPlayer("newPlayer", c)
}

func (m *GameManager) broadcastLeavingPlayer( c *GameClient) {
	m.broadcastPlayer("leavingPlayer", c)
}
