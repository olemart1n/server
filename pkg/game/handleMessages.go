package game

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func (c *GameClient) handleMessages() {
	defer func() {
		c.gameManager.removeGameClient(c)
	}()

	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}

	c.connection.SetReadLimit(512)
	c.connection.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			log.Println("error in handleMessages.go: ", err)
			break // Break the loop to close conn & Cleanup
		}

		var request GameEvent
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error when marshalling message: %v", err)
			break // Breaking the connection here might be harsh
		}
		for client := range c.gameManager.gameClients {
			if c != client {
				client.egress <- request
			}
	
		}
	}
}