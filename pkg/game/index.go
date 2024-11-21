package game

func NewManager() *Manager {
	m := &Manager{
		Players: make(Players),
		Spectators: make(Spectators),
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
	for player := range m.Spectators {
		p := PlayerClientData{}
		p.Username =player.Username
		p.ID = player.Id
		players = append(players, p)
	}
	return players
}
