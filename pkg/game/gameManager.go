package game

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)


var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
				origin := r.Header.Get("Origin")
				switch origin {
				case "http://localhost:5173":
					return true
				case "https://olems.no":
					return true
				}
				return false
			},
		EnableCompression: true,
	}
)

type GameClientList map[*GameClient]bool

type GameManager struct {
	gameClients GameClientList
	sync.RWMutex
}

func NewGameManager() *GameManager {
	m := &GameManager{
		gameClients: make(GameClientList),
	}
	return m
}

func (m *GameManager) ServeGameWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	username := r.URL.Query().Get("username")
	gameClient := NewGameClient(conn, m, username)
	m.addGameClient(gameClient)
	go gameClient.handleMessages()
	go gameClient.sendMessages()
}


func (m * GameManager) addGameClient (c *GameClient) {
	m.Lock()
	defer m.Unlock()
	m.gameClients[c] = true
	sendConnectedPlayersLength(c)
	c.sendExistingPlayersToNewGameClient(m)
	m.broadcastNewPlayer(c)
}

func (m * GameManager) removeGameClient (c *GameClient) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.gameClients[c]; ok {
		c.connection.Close()
		delete(m.gameClients, c)
		m.broadcastLeavingPlayer(c)
	}
	
}