package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {

	// Create a new router instance with the default settings
	router := gin.Default()

	// Initialize and register all the routes
	initializeRoutes(router)

	// Run the router on port 8080
	router.Run(":8080") // listen and serve on

}
