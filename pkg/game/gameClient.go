package game

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)



type GameEvent struct {
	GameEvent string `json:"name"`
	Payload json.RawMessage `json:"payload"`
}

type GameClient struct {
	connection *websocket.Conn
	gameManager *GameManager
	egress chan GameEvent
	username string
}

func NewGameClient(conn *websocket.Conn, manager *GameManager, username string) *GameClient {
	return &GameClient{
		username: username,
		connection: conn,
		gameManager:    manager,
		egress: make(chan GameEvent),
	}
}


func (c *GameClient) sendExistingPlayersToNewGameClient (m *GameManager) {
	// Step 1: Create a list of all existing players except the new one
	existingPlayers := []string{}
	for player := range m.gameClients {
		if player != c {
			existingPlayers = append(existingPlayers, player.username)
		}
	}


    // Step 2: Send the list of existing players to the new player
    dataForNewPlayer := map[string]interface{}{
        "name": "existingPlayers",
        "payload": map[string]interface{}{
            "players": existingPlayers,
        },
    }
    jsonDataForNewPlayer, err := json.Marshal(dataForNewPlayer)
    if err != nil {
        log.Println(err)
        return
    }
    if err := c.connection.WriteMessage(websocket.TextMessage, jsonDataForNewPlayer); err != nil {
        log.Println(err)
        return
    }
}