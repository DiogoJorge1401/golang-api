package main

import (
	"github.com/DiogoJorge1401/golang-api/config"
	"github.com/DiogoJorge1401/golang-api/router"
)

func main() {
	err := config.InitializeDatabaseConnection()
	logger := config.GetLoger()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.InitializeRouter()
}
