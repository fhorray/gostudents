package config

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "simpleapi"
)


var (
	db 			*gorm.DB
	logger 	*Logger
)

func Init() error {
	var err error

	// Initialize SqLite
	db, err = InitializeDB()

	if err != nil {
		return fmt.Errorf("Error initializing SqLite: %v", err)
	}
	return nil
}

func GetSqlite() *gorm.DB {
	return db
}


func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}

