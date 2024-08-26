package chat

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type Event struct {
	Name string `json:"name"`
	Payload json.RawMessage `json:"payload"`
}




type Client struct {
	connection *websocket.Conn
	manager *Manager
	egress chan Event
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress: make(chan Event),
	}
}

