package utils

import (
	"encoding/json"
	"fmt"

	"github.com/olemart1n/server/pkg/game/event"
)

func TypeAsserter[T any](request event.Message) (T, error) {
    var result T
    payloadBytes, err := json.Marshal(request.Payload) // Convert interface{} to JSON bytes
    if err != nil {
        return result, fmt.Errorf("failed to marshal payload: %w", err)
    }

    err = json.Unmarshal(payloadBytes, &result) // Convert JSON bytes to the target type
    if err != nil {
        return result, fmt.Errorf("failed to unmarshal payload: %w", err)
    }

    return result, nil
}