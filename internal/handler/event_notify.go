package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gpt-chat/internal/chat"
	"gpt-chat/internal/meta"
	"log"
)

type EventNotify struct {
	webService *meta.Service
	box        *chat.Box
}

func NewEventNotify(webService *meta.Service, chatBox *chat.Box) *EventNotify {
	return &EventNotify{
		webService: webService,
		box:        chatBox,
	}
}

func (n EventNotify) Handle(ctx *gin.Context) {
	data, err := n.contextRawData(ctx)
	if err != nil {
		log.Printf("Error while getting request's raw data: %s", err)
		return
	}
	var event meta.Event
	err = json.Unmarshal(data, &event)
	if err != nil {
		log.Printf("Cannot unmarshal event data: %s", err)
		return
	}
	if event.TextBody() == "" || event.WhatsappId() == "" {
		return
	}
	log.Printf(
		"Received notify about message: '%s' from %s (%s)",
		event.TextBody(),
		event.ProfileName(),
		event.WhatsappId())

	var replyMessage *chat.Message
	replyMessage, err = n.box.ResponseOn(&chat.Message{
		Handle: event.WhatsappId(),
		Text:   event.TextBody(),
	})
	if err != nil {
		log.Printf("Error while getting response from chatbox: %s", err)
		return
	}
	log.Printf("Reply with: %s", replyMessage.Text)
	err = n.sendMessage(event.WhatsappId(), replyMessage.Text)
	if err != nil {
		log.Printf("Error while sending message to Whatsapp ID %s: %s", event.WhatsappId(), err)
	}
}

func (n EventNotify) contextRawData(ctx *gin.Context) ([]byte, error) {
	if ctx.ContentType() != "application/json" {
		return nil, fmt.Errorf("invalid content type '%s', expected 'application/json'", ctx.ContentType())
	}
	data, err := ctx.GetRawData()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (n EventNotify) sendMessage(to string, textBody string) error {
	request := meta.TextMessageRequest{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               to,
		Type:             "text",
		Text: meta.TextMessageText{
			PreviewUrl: false,
			Body:       textBody,
		},
	}
	response, err := n.webService.SendRequest(request)
	if !response.IsHttpStatusSuccess() {
		return fmt.Errorf(
			"web service response (%d) is not successful: %s",
			response.HttpResponse.StatusCode,
			response.Body())
	}

	return err
}
