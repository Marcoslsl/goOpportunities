package config

import (
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB
var logger *Logger

func Init() error {
	var err error

	// INititalize sqlite
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initialize sqlite %v", err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize logger
	logger = NewLogger(p)
	return logger
}
