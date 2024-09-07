package config

import (
	"os"

	"github.com/DiogoJorge1401/golang-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLoger("sqlite")
	databasePath := "./db/main.db"

	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		if err := os.MkdirAll("./db", os.ModePerm); err != nil {
			return nil, err
		}

		if file, err := os.Create(databasePath); err != nil {
			return nil, err
		} else {
			file.Close()
		}
	}

	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})

	if err != nil {
		logger.Errorf("error when connecting to database: %v", err)
		return nil, err
	}

	if err = db.AutoMigrate(&schemas.Oppening{}); err != nil {
		logger.Errorf("error when migrating schemas: %v", err)
		return nil, err
	}

	return db, nil
}
