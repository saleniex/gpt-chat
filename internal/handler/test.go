package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Test struct {
}

func (w Test) Handle(ctx *gin.Context) {
	log.Println("Content type: " + ctx.ContentType())
	data, err := ctx.GetRawData()
	if err != nil {
		log.Println("Error: ", err)
	}
	log.Println(string(data))
}
