package game

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/olemart1n/server/pkg/game/event"
)


type Players map[*Client]bool
type Spectators map[*Client]bool

type Manager struct {
	Players Players
	Spectators Spectators
	sync.RWMutex
}



type Client struct {
	Connection *websocket.Conn
	Egress chan event.Message
	Username string
	Id string
	Hp int // 0 means spectator, > 0 means player
}


type PlayerClientData struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

type DamageData struct {
	VictimId string `json:"victimId"`
	Amount   int    `json:"amount"`
	ShooterId   int    `json:"shooterId"`
}