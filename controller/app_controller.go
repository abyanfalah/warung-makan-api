package controller

import (
	"net/http"
	"warung-makan/config"
	"warung-makan/manager"
	"warung-makan/model"
	"warung-makan/utils"
	"warung-makan/utils/authenticator"

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

	test := router.Group("/test")

	// GENERATE TOKEN
	test.POST("/generate_token", func(ctx *gin.Context) {
		var user model.User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			utils.JsonErrorBadRequest(ctx, err, "cannot bind struct")
		}

		accessToken := authenticator.NewAccessToken(config.NewConfig().TokenConfig)
		token, err := accessToken.GenerateAccessToken(&user)
		if err != nil {
			utils.JsonErrorBadGateway(ctx, err, "cannot generate token")
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})

	return &controller
}
