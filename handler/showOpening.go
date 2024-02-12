package handler

import (
	"fmt"
	"github.com/batidou/gopportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowOpeningHandler(ctx *gin.Context) {

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	// find opening by id

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSSucess(ctx, "show-opening", opening)
}
