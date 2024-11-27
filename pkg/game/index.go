package game

import (
	"database/sql"
	"log"
	"os"

	"github.com/olemart1n/server/pkg/game/turso"
)

func NewManager() *Manager {
	var db *sql.DB
	var err error

	if os.Getenv("PRODUCTION") == "true" {
		db, err = turso.ProdTursoDB()
	} else {
		log.Println("DEV DB")
		db, err = turso.DevTursoDB()
	}
	if err != nil {
		log.Print(err)
	}

	m := &Manager{
		Players: make(Players),
		Spectators: make(Spectators),
		DB: db,
	}
	return m
}


func getSpectatorList (m *Manager) []PlayerClientData {
	spectators  := []PlayerClientData{}
	for spectator := range m.Spectators {
		p := PlayerClientData{}
		p.Username =spectator.Username
		p.ID = spectator.Id
		spectators = append(spectators, p)
	}
	return spectators
}
func getPlayerList (m *Manager) []PlayerClientData {
	players  := []PlayerClientData{}
	for player := range m.Players {
		p := PlayerClientData{}
		p.Username =player.Username
		p.ID = player.Id
		players = append(players, p)
	}
	return players
}
