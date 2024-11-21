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
	}

	return client, id
}
