package chat

import (
	"time"
)

// pongHandler is used to handle PongMessages for the Client
func (c *Client) pongHandler(pongMsg string) error {
	// Current time + Pong Wait time
	
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}

