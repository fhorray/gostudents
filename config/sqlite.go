package config

import (
	"os"
	"studentsapi/schemas"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if database exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")

		// Create database and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create and connect
	db, err :=	gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errf("Sqlite opening error")
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Students)
	if err != nil {
		logger.Errf("Sqlite automigration error")
		return nil, err
	}
	// Return database
	return db, nil
}
