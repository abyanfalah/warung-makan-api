package controller

import (
	"net/http"
	"warung-makan/config"
	"warung-makan/model"
	"warung-makan/usecase"
	"warung-makan/utils"
	"warung-makan/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	usecase usecase.UserUsecase
	router  *gin.Engine
}

func (lc *LoginController) Login(ctx *gin.Context) {
	var credential model.Credential
	accessToken := authenticator.NewAccessToken(config.NewConfig().TokenConfig)

	err := ctx.ShouldBindJSON(&credential)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	user, err := lc.usecase.GetByCredentials(credential.Username, credential.Password)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "invalid credentials")
		return
	}

	token, err := accessToken.GenerateAccessToken(&user)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "cannot generate token")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "you are logged in",
		"token":   token,
	})
}

func (lc *LoginController) LoginTest(ctx *gin.Context) {
	var credential model.Credential
	accessToken := authenticator.NewAccessToken(config.NewConfig().TokenConfig)

	err := ctx.ShouldBindJSON(&credential)

	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	user, err := lc.usecase.GetByCredentials(credential.Username, credential.Password)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "invalid credentials")
		return
	}

	_, err = accessToken.GenerateAccessToken(&user)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "cannot generate token")
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func NewLoginController(usecase usecase.UserUsecase, router *gin.Engine) *LoginController {
	controller := LoginController{
		usecase: usecase,
		router:  router,
	}

	router.POST("/login", controller.Login)
	router.POST("/test/login", controller.LoginTest)

	return &controller
}
