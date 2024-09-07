package router

import (
	"github.com/DiogoJorge1401/golang-api/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	handler.Init()

	InitializeOpeningRoutes(v1)
}
