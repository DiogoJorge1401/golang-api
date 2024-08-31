package router

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	InitializeOppeningRoutes(v1)
}
