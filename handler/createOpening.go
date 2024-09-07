package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := new(struct {
		Role string `json:"role"`
	})

	ctx.BindJSON(request)

	logger.Infof("request body received: %+v", *request)
}
