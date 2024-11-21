package utils

import (
	"net/http"

	"github.com/gorilla/websocket"
)


var (
	WebsocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			switch origin {
			case "http://localhost:5173":
				return true
			case "https://olems.no":
				return true
			case "https://www.olems.no":
				return true
			}
			return false
		},
		EnableCompression: true,
	}
)