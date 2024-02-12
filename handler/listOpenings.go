package handler

import (
	"fmt"
	"github.com/batidou/gopportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("error listing openings: %v", err.Error()))
		return
	}

	sendSSucess(ctx, "list-openings", openings)
}
