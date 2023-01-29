package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Root struct{}

func NewRoot() *Root {
	return &Root{}
}

func (r Root) Handle(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
