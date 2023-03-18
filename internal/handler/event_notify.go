package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gpt-chat/internal/meta"
	"log"
)

type EventNotify struct {
}

func (n EventNotify) Handle(ctx *gin.Context) {
	log.Println("Content type: " + ctx.ContentType())
	data, err := ctx.GetRawData()
	if err != nil {
		log.Printf("Error while retting request's raw data: %s", err)
		return
	}
	var event meta.Event
	err = json.Unmarshal(data, &event)
	if err != nil {
		log.Printf("Cannot unmarshar event data: %s", err)
		return
	}
	log.Printf("Received: '%s' from %s (%s)", event.TextBody(), event.ProfileName(), event.WhatsappId())
}
