package router

import (
	"github.com/DiogoJorge1401/golang-api/handler"
	"github.com/gin-gonic/gin"
)

func InitializeOppeningRoutes(router *gin.RouterGroup) {
	oppeningsRouter := router.Group("/oppenings")

	oppeningsRouter.GET("/", handler.ListOppeningsHandler)
	oppeningsRouter.GET("/:id", handler.FindOppeningHandler)
	oppeningsRouter.POST("/", handler.CreateOppeningHandler)
	oppeningsRouter.DELETE("/:id", handler.DeleteOppeningHandler)
	oppeningsRouter.PUT("/:id", handler.UpdateOppeningHandler)
}
