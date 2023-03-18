package meta

import (
	"encoding/json"
	"os"
	"testing"
)

func TestMetaEventStructure(t *testing.T) {
	bytes, err := os.ReadFile("event.json")
	if err != nil {
		t.Errorf("Cannot read event JSON test object from file")
		return
	}

	var event Event
	err = json.Unmarshal(bytes, &event)
	if err != nil {
		t.Errorf("Cannot unmarshal test JSON object: %s", err)
	}

	if event.TextBody() != "_TEXT_BODY_" {
		t.Errorf("Invalid unmashalled content. Unexpected text body '%s'.", event.TextBody())
	}
	if event.ProfileName() != "Uldis Jansons" {
		t.Errorf("Invalid unmashalled content. Unexpected profile name '%s'.", event.ProfileName())
	}
	if event.WhatsappId() != "37120042219" {
		t.Errorf("Invalid unmashalled content. Unexpected whatsapp ID '%s'", event.WhatsappId())
	}

}
