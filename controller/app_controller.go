package controller

import (
	"net/http"
	"warung-makan/manager"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ucMan  manager.UsecaseManager
	router *gin.Engine
}

func NewController(usecaseManager manager.UsecaseManager, router *gin.Engine) *Controller {
	controller := Controller{
		ucMan:  usecaseManager,
		router: router,
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	return &controller
}
