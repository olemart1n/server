package game

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)


func NewGameClient(conn *websocket.Conn, manager *GameManager, username string) *GameClient {
	id := uuid.New().String()
	client := &GameClient{
		username: username,
		connection: conn,
		gameManager:    manager,
		egress: make(chan Message),
		id: id,
	}

	jsonData, err := createJsonObject("id", id)
	if err != nil {
		print(err)
	}

	client.connection.WriteMessage(websocket.TextMessage, jsonData)

	return client
}


func (c *GameClient) sendConnectedPlayers () error {
	playerList := getPlayerList(c.gameManager)
    jsonData, err := createJsonObject("connected_players", playerList)
    if err != nil {
		print(err)
        return err
    }
    if err := c.connection.WriteMessage(websocket.TextMessage, jsonData); err != nil {
        return err
    }
	return nil
}