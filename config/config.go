package config

import (
	"gorm.io/gorm"
)

var (
	// DBConfig is the configuration for the database
	bd     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
