package game

import (
	"context"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/olemart1n/server/pkg/game/utils"
)



func (c *Client) handlePingPong(m *Manager, ctx context.Context, cancel context.CancelFunc) {
    ticker := time.NewTicker(utils.PingInterval) // For sending PING messages
    defer func() {
        ticker.Stop()
        m.removeSpectator(c)
    }()

    // Set the PONG handler to reset the read deadline
    c.Connection.SetPongHandler(func(appData string) error {
        return c.Connection.SetReadDeadline(time.Now().Add(utils.PongWait))
    })

    // Initial read deadline
    c.Connection.SetReadDeadline(time.Now().Add(utils.PongWait))

    for {
        select {
        case <-ctx.Done(): // Stop if context is canceled
            return
        case <-ticker.C:
			spectators := getSpectatorList(m)

			data, _ := utils.CreateJsonObject("connected_spectators", spectators)

			if err := c.Connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println("Error sending connected spectators")
			}
            
            if err := c.Connection.WriteMessage(websocket.PingMessage, nil); err != nil {
                log.Println("Error sending PING:", err)
                cancel() // Signal all routines to stop
                return
            }
        }
    }
}