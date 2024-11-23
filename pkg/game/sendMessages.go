package game

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

func (c *Client) sendMessages(ctx context.Context, cancel context.CancelFunc) {


	for {
		select {
		case <- ctx.Done():
			return
		case event, ok := <-c.Egress:
			if !ok {
				if err := c.Connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Printf("error in (case message, ok := <- c.egress): %s",err)
				}
				cancel()
				return
			}
			data, err := json.Marshal(event)
			if err != nil {
				log.Println(err)
				cancel()
				return
			}

			if err := c.Connection.WriteMessage(websocket.TextMessage, data); err != nil{
				log.Println(err)
				return
			}

		}
	}
}
