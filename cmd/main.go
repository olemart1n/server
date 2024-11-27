package main

import (
	"log"
	"net/http"
	"os"

	"github.com/olemart1n/server/pkg/chat"
	"github.com/olemart1n/server/pkg/game"
	"github.com/olemart1n/server/pkg/handlelista"
)

func main() {
	mistralClient := handlelista.NewMistralClient()
	manager := chat.NewManager()
	gameManager := game.NewManager()

	http.HandleFunc("/car-game", gameManager.ServeGameWS)
	http.HandleFunc("/car-game-players", gameManager.SendPlayersViaHTTP)
	http.HandleFunc("/ws",manager.ServeWS)
    http.HandleFunc("/robokokk/prompt1", func(w http.ResponseWriter, r *http.Request) {
        handlelista.Prompt(mistralClient, w, r)
    })


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
