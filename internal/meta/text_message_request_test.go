package meta

import (
	"testing"
)

func TestTextMessageRequest_Data(t *testing.T) {
	req := TextMessageRequest{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               "37120042219",
		Type:             "text",
		Text: TextMessageText{
			PreviewUrl: false,
			Body:       "Test message",
		},
	}
	data, err := req.Data()
	if err != nil {
		t.Errorf("Error while getting data: %s", err)
	}
	if data != "{\"messaging_product\":\"whatsapp\",\"recipient_type\":\"individual\",\"to\":\"37120042219\",\"type\":\"text\",\"text\":{\"preview_url\":false,\"body\":\"Test message\"}}" {
		t.Errorf("Data does not match")
	}
}
