package router

import (
	"github.com/DiogoJorge1401/golang-api/handler"
	"github.com/gin-gonic/gin"
)

func InitializeOpeningRoutes(router *gin.RouterGroup) {
	openingsRouter := router.Group("/openings")

	openingsRouter.GET("/", handler.ListOpeningsHandler)
	openingsRouter.GET("/:id", handler.FindOpeningHandler)
	openingsRouter.POST("/", handler.CreateOpeningHandler)
	openingsRouter.DELETE("/:id", handler.DeleteOpeningHandler)
	openingsRouter.PUT("/:id", handler.UpdateOpeningHandler)
}
