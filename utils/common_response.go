package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonDataResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func JsonDataMessageResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": message,
	})
}

func JsonNamedDataMessageResponse(ctx *gin.Context, keyName string, data interface{}, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		keyName:   data,
		"message": message,
	})
}

func JsonSuccessMessage(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func JsonBadRequestMessage(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}

func JsonBadGatewayMessage(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
		"message": message,
	})
}

func JsonErrorBadGateway(ctx *gin.Context, err error, message string) {
	ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
		"error":   err.Error(),
		"message": message,
	})
}

func JsonErrorBadRequest(ctx *gin.Context, err error, message string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error":   err.Error(),
		"message": message,
	})
}

func JsonErrorUnauthorized(ctx *gin.Context, err error, message string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error":   err.Error(),
		"message": message,
	})
}
