package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	// Create a new router instance with the default settings
	router := gin.Default()

	// Define a route for the path /ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on

}