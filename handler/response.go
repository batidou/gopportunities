package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"Message":   message,
		"errorCode": code,
	})
}

func sendSSucess(ctx *gin.Context, operation string, data interface{}) {

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, gin.H{
		"message": fmt.Sprintf("operation from handler: %s was successful", operation),
		"data":    data,
	})
}
