package chat

// Message with conversation handle
// Used in REST API message (JSON object) marshalling and unmarshalling
type Message struct {
	Handle string `json:"handle"`
	Text   string `json:"text"`
}
