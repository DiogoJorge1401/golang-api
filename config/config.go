package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		return err
	}

	return err

}

func GetLoger(prefix string) *Logger {
	if logger == nil {
		logger = newLoger(prefix)
	}

	return logger
}
