package game

import (
	"time"
)


var (
	pongWait = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

func (c *GameClient) pongHandler(pongMsg string) error {
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}