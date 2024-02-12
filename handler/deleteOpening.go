package handler

import (
	"fmt"
	"github.com/batidou/gopportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

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

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id, err.Error()))
		return
	}

	sendSSucess(ctx, "delete-opening", opening)
}
