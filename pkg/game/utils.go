package game

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

func (m *GameManager) broadcastPlayer(name string, c *GameClient) {

	data, err := createJsonObject(name, PlayerClientData{Username: c.username, ID: c.id})
	
	if err != nil {
		log.Println(err)
		return
	}

	for player := range m.gameClients {
		if player != c {
			if err := player.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
				return
			}
		}	
	}
}

func (m *GameManager) broadcastNewPlayer( c *GameClient) {
	m.broadcastPlayer("new_player", c)
}

func (m *GameManager) broadcastLeavingPlayer( c *GameClient) {
	m.broadcastPlayer("leaving_player", c)
}


// func getUsernames (m *GameManager) []string {
// 	usernames := []string{}
// 	for player := range m.gameClients {
// 		usernames = append(usernames, player.username)
// 	}
// 	return usernames
// }

func getPlayerList (m *GameManager) []PlayerClientData {
	playerList  := []PlayerClientData{}
	for player := range m.gameClients {
		p := PlayerClientData{}
		p.Username = player.username
		p.ID = player.id
		playerList = append(playerList, p)
	}
	return playerList
}

func createJsonObject(name string, data interface{}) ([]byte, error) {
    message := Message{
        Message: name,
        Payload: data,
    }
    jsonData, err := json.Marshal(message)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    return jsonData, nil
}