package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type VerifyToken struct {
	Token string
}

func (t VerifyToken) Handle(ctx *gin.Context) {
	hubMode := ctx.Query("hub.mode")
	if hubMode != "subscribe" {
		log.Println("Cannot verify token. Invalid mode.")
		ctx.Status(http.StatusBadRequest)
		return
	}

	verifyToken := ctx.Query("hub.verify_token")
	if verifyToken != t.Token {
		log.Println("Cannot verify token. Provided invalid token.")
		ctx.Status(http.StatusBadRequest)
		return
	}

	hubChallenge := ctx.Query("hub.challenge")
	ctx.Data(http.StatusOK, "text/html", []byte(hubChallenge))
}
