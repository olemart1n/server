package handlelista

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gage-technologies/mistral-go"
)



func  Prompt (c *mistral.MistralClient, w http.ResponseWriter, r *http.Request) {
    recommendation, err := mistralChat(c, createPrompt1())
    w.Header().Set("Content-Type", "application/json")

    if err != nil {
        w.WriteHeader(http.StatusServiceUnavailable)
        json.NewEncoder(w).Encode(map[string]string{"error": "mistral error"})
        return
    }
    json.NewEncoder(w).Encode(recommendation)
}




func createPrompt1 () (mistral.ChatMessage) {
    currentTime := time.Now().Format(time.RFC3339)
    content := currentTime + promptOne
    createPrompt1 := mistral.ChatMessage {
		Role: "user",
		Content: content,
	}
    return createPrompt1
}

func  mistralChat(m *mistral.MistralClient, prompt mistral.ChatMessage) (string, error) {
    response, err := m.Chat("mistral-tiny", []mistral.ChatMessage{prompt}, nil)
    if err != nil {
        fmt.Print(err)
        return "", err
    }
    if response.Choices[0].Message.Content == "" {
        fmt.Println("Empty response received from Mistral")
    }
    return response.Choices[0].Message.Content, nil
}