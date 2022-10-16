package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonErrorResponse(ctx *gin.Context, err error, message string) {
	ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
		"error":   err,
		"message": message,
	})
}
