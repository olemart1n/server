package game

import "net/http"


func (m *GameManager) SendPlayersViaHTTP(w http.ResponseWriter, r *http.Request) {
    // Handle CORS headers
    origin := r.Header.Get("Origin")
    switch origin {
    case "http://localhost:5173", "https://olems.no", "https://www.olems.no":
        w.Header().Set("Access-Control-Allow-Origin", origin)
        w.Header().Set("Access-Control-Allow-Methods", "GET")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    default:
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    // Prepare response data (example logic)

    players := getPlayerList(m)

	data, err := createJsonObject("connected_players", players)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
}