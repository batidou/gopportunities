package router

import (
	"github.com/batidou/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	// initialize handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		const OpeningPath = "/opening"
		const OpeningsPath = "/openings"

		v1.GET(OpeningPath, handler.ShowOpeningHandler)
		v1.POST(OpeningPath, handler.CreateOpeningHandler)
		v1.DELETE(OpeningPath, handler.DeleteOpeningHandler)
		v1.PUT(OpeningPath, handler.UpdateOpeningHandler)
		v1.GET(OpeningsPath, handler.ListOpeningsHandler)

	}
}
