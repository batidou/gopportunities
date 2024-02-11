package main

import (
	"github.com/batidou/gopportunities/config"
	"github.com/batidou/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// Initialize configuration
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configuration: %v", err)
		return

	}

	// initialize the router
	router.InitializeRouter()

}
