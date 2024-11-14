package game

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func (c *GameClient) sendMessages() {

	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		c.gameManager.removeGameClient(c)
	}()

	for {
		select {
		case event, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Printf("error in (case message, ok := <- c.egress): %s",err)
				}
				return
			}
			data, err := json.Marshal(event)
			if err != nil {
				log.Println(err)
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil{
				log.Println(err)
				return
			}
		case <- ticker.C:
			if err := c.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println(err)
				return
			}
		}
	}
}