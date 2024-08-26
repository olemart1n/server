package chat

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)



func (c *Client) readAndBroadCast() {
	defer func() {
		c.manager.removeClient(c)
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
				log.Printf("error reading message in readMessages.go: %v", err)
			}
			log.Println("error in readMessages.go: ", err)
			break // Break the loop to close conn & Cleanup
		}
		var request Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error marshalling message in readMessages.go: %v", err)
			break // Breaking the connection here might be harsh xD
		}

		for client := range c.manager.clients {
			if c != client {
				client.egress <- request
			}
	
		}
		
	}
}