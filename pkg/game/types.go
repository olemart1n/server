package game

import (
	"sync"

	"github.com/gorilla/websocket"
)



type Message struct {
	Message string `json:"name"`
	Payload interface{} `json:"payload"`
}


type GameClientList map[*GameClient]bool

type GameManager struct {
	gameClients GameClientList
	sync.RWMutex
}



type GameClient struct {
	connection *websocket.Conn
	gameManager *GameManager
	egress chan Message
	username string
	id string
}


type PlayerClientData struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}