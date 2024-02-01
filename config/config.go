package config

import (
	"gorm.io/driver/postgres"
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
	db *gorm.DB
	logger *Logger
)

func Init() (*gorm.DB, error) {
	dsn := "user=seu_usuario password=sua_senha dbname=seu_banco_de_dados sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}

