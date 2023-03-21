package meta

import "encoding/json"

// TextMessageRequest Implementation of meta.ServiceRequest interface
//
// Request used to send message to Whatsapp user
type TextMessageRequest struct {
	MessagingProduct string          `json:"messaging_product"`
	RecipientType    string          `json:"recipient_type"`
	To               string          `json:"to"`
	Type             string          `json:"type"`
	Text             TextMessageText `json:"text"`
}

type TextMessageText struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

func (r TextMessageRequest) Data() (string, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(data), err
}
