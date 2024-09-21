package router

import (
	"github.com/DiogoJorge1401/golang-api/middleware"
	"github.com/gin-gonic/gin"

	docs "github.com/DiogoJorge1401/golang-api/docs"

	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRouter() {
	var router = gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	router.Use(middleware.MalformedBodyMiddleware)

	initializeRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
