package game

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/olemart1n/server/pkg/game/event"
)

func NewSpectator (conn *websocket.Conn, username string) (*Client, string) {
	id := uuid.New().String()
	client := &Client{
		Username: username,
		Connection: conn,
		Egress: make(chan event.Message),
		Id: id,
		Hp: 100,
	}

	return client, id
}


func findPlayerById(players Players, id string) *Client {
	for client := range players {
		if client.Id == id {
			return client
		}
	}
	return nil
}