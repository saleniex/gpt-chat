package handler

import (
	"chat-gpt-rest/internal/chat"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Say struct {
	box *chat.Box
}

func NewSay(box *chat.Box) *Say {
	return &Say{
		box,
	}
}

func (s *Say) Handle(ctx *gin.Context) {
	var message *chat.Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Cannot bind params " + err.Error(),
		})
		return
	}
	responseMessage, err := s.box.ResponseOn(message)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Error while getting response from box: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responseMessage)
}
