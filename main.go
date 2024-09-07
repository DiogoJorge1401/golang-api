package main

import (
	"github.com/DiogoJorge1401/golang-api/config"
	"github.com/DiogoJorge1401/golang-api/router"
)

func main() {
	logger := config.GetLoger("main")

	if err := config.Init(); err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.InitializeRouter()
}
