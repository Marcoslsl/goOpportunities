package main

import (
	"github.com/Marcoslsl/goOpportunities.git/config"
	r "github.com/Marcoslsl/goOpportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	// Initialize Cconfigs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize router
	r.Initialize()
}
