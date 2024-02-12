package handler

import (
	"github.com/batidou/gopportunities/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	ctx.Header("Content-Type", "application/json")

	if err := request.Validate(); err != nil {
		logger.Errorf("request validation failed: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSSucess(ctx, "create-opening", opening)
}
