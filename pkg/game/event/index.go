package event


type Message struct {
	Message string `json:"name"`
	Payload interface{} `json:"payload"`
}

