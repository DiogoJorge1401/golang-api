package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	var router = gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	router.Run(":8080")
}
