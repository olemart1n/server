package game

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/olemart1n/server/pkg/game/utils"
)

func (m *Manager) broadcastPlayerToAll(name string, c *Client, filterOutClient bool) {

	data, err := utils.CreateJsonObject(name, PlayerClientData{Username: c.Username, ID: c.Id})	
	if err != nil {
		log.Println(err)
		return
	}

	if filterOutClient {
		for spectator := range m.Spectators {
			if c != spectator {
				if err := spectator.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
					log.Println(err)
					return
				}
			}
	}
	} else {
		for spectator := range m.Spectators {
			if err := spectator.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
				return
			}	
		}
	}
}

func (m *Manager) sendAlreadyActivePlayers(c *Client) {
	data, err := utils.CreateJsonObject("already_active_players", getPlayerList(m))	
	if err != nil {
		log.Println("error when creating json object of getPlayerList()")
		log.Println(err)
	}
	if err := c.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Println(err)
		return
	}
}


// func (m *Manager) broadcastNewPlayer( c *Client) {
// 	m.broadcastMessageToAll("spectator_joins", c)
// }

// func (m *Manager) broadcastLeavingPlayer( c *Client) {
// 	m.broadcastMessageToAll("spectator_leaves", c)
// }
