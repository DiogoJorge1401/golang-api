package handler

import (
	"github.com/DiogoJorge1401/golang-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	var err error
	logger = config.GetLoger("handler")
	db, err = config.GetDB()

	if err != nil {
		logger.Errorf("error when connecting to db: %v", err)
	}
}
