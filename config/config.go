package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func GetDB() (*gorm.DB, error) {
	var err error

	if db != nil {
		return db, nil
	}

	db, err = InitializeSQLite()

	if err != nil {
		return nil, err
	}

	return db, err

}

func GetLoger(prefix string) *Logger {
	if logger == nil {
		logger = newLoger(prefix)
	}

	return logger
}
