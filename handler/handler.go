package handler

import (
	"github.com/Marcoslsl/goOpportunities.git/config"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
