package game

import (
	"log"

	"github.com/olemart1n/server/pkg/game/event"
)



func toSpectatorEgressToAll(m *Manager, request event.Message) {
	for spectator := range m.Spectators {
		
			select {
			case spectator.Egress <- request:
				// Successfully sent
			default:
				log.Printf("Dropping message for spectator %s due to full buffer", spectator.Id)
			}
		
	}
}

func toSpectatorEgressFilterOutClient(c *Client, m *Manager, request event.Message) {
	for spectator := range m.Spectators {
		if c != spectator {
			select {
			case spectator.Egress <- request:
				// Successfully sent
			default:
				log.Printf("Dropping message for spectator %s due to full buffer", spectator.Id)
			}
		}
	}
}
func toPlayerEgressToAll(m *Manager, request event.Message) {
	for player := range m.Players {
		
			select {
			case player.Egress <- request:
				// Successfully sent
			default:
				log.Printf("Dropping message for spectator %s due to full buffer", player.Id)
			}
		
	}
}

func toPlayerEgressFilterOutClient(c *Client, m *Manager, request event.Message) {
	for player := range m.Players {
		if c != player {
			select {
			case player.Egress <- request:
				// Successfully sent
			default:
				log.Printf("Dropping message for spectator %s due to full buffer", player.Id)
			}
		}
	}
}