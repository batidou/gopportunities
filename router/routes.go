package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		const OpeningPath = "/opening"

		v1.GET(OpeningPath, func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET Opening",
			})
		})

		v1.POST(OpeningPath, func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "POST Opening",
			})
		})

		v1.DELETE(OpeningPath, func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "DELETE Opening",
			})
		})

		v1.PUT(OpeningPath, func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "PUT Opening",
			})
		})

		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET Openings",
			})
		})
	}
}
