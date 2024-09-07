package router

import (
	"github.com/DiogoJorge1401/golang-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	var router = gin.Default()

	router.Use(middleware.MalformedBodyMiddleware)

	initializeRoutes(router)

	router.Run(":8080")
}
