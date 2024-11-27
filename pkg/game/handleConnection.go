package game

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/olemart1n/server/pkg/game/event"
	"github.com/olemart1n/server/pkg/game/turso"
	"github.com/olemart1n/server/pkg/game/utils"
)

func (c *Client) handleConnection(m *Manager, cancel context.CancelFunc) {

	defer func() {
        m.removeSpectator(c) // Remove client from the Spectator map
		cancel()
    }()
	
	for {		
		_, payload, err := c.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error: %v", err)
			} else if errors.Is(err, os.ErrDeadlineExceeded) {
				log.Println("ReadMessage timeout: no data received within PongWait")
			}
			log.Println("Exiting handleConnection due to error: ", err)
			break
		}
		

		var request event.Message
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error when marshalling message: %v", err)
			continue
		}

		switch request.Message  {
		case "player_joins":
			m.Lock()
            if _, playing := m.Players[c]; !playing {
				m.sendAlreadyActivePlayers(c)
				c.Hp = 100
                m.Players[c] = true // Add to playingClients
				m.broadcastPlayerToAll("player_joins", c, false)
				
            }
            m.Unlock()
		case "chat_message":
            m.Lock()
			toSpectatorEgressFilterOutClient(c, m, request)

			data,_ := utils.TypeAsserter[ChatMessage](request)
			err := turso.InsertMessage(m.DB, data.SenderUsername, data.Message, data.SenderId);if err!= nil {
				log.Print(err)
			}
            m.Unlock()
		case "hp_damage":
			m.RLock() // Read lock because we're just iterating over Players

			data, _ := utils.TypeAsserter[DamageData](request)

			victim :=findPlayerById(m.Players, data.VictimId)


			victim.Hp -= data.Damage
			if(victim.Hp <= 0) {
				delete(m.Players, victim)
				m.broadcastPlayerToAll("player_died", victim, false)
			} else {
				toPlayerEgressToAll(m, request)
			}
			m.RUnlock()
		case "car_data":
			toPlayerEgressFilterOutClient(c, m, request)

		default:
			m.RLock()
			toPlayerEgressFilterOutClient(c, m, request)
			m.RUnlock()
		}
	}
}