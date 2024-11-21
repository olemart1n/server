package utils

import (
	"encoding/json"
	"log"

	"github.com/olemart1n/server/pkg/game/event"
)



func CreateJsonObject(name string, data interface{}) ([]byte, error) {
    message := event.Message{
        Message: name,
        Payload: data,
    }
    jsonData, err := json.Marshal(message)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    return jsonData, nil
}