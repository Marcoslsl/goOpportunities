package config

import (
	"errors"

	"gorm.io/gorm"
)

var db *gorm.DB
var logger *Logger

func Init() error {
	return errors.New("fake error")
	// return nill
}

func GetLogger(p string) *Logger {
	// Initialize logger
	logger = NewLogger(p)
	return logger
}
