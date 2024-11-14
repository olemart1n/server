package game

import (
	"log"
	"net/http"
)


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
	c.sendConnectedPlayers()
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