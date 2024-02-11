package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	// DBConfig is the configuration for the database
	bd     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	// Initialize the MySQL database
	bd, err = InitializeMySQL()
	if err != nil {
		return fmt.Errorf("Error initializing MySQL: %v", err)
	}
	return nil
}

func GetMySQLDB() *gorm.DB {
	return bd
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
