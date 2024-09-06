package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitializeDatabaseConnection() error {
	return nil
}

func GetLoger() *Logger {
	if logger == nil {
		logger = newLoger("")
	}

	return logger
}
