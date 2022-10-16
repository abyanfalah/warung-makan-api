package controller

import (
	"net/http"

	"warung-makan/manager"
	"warung-makan/model"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	ucMan  UsecaseManager
	router *gin.Engine
}

func (c *UserController) CreateNewUser(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func NewUserController(usecaseManager manager.UsecaseManager, router *gin.Engine) *UserController {
	controller := UserController{
		ucMan:  usecaseManager,
		router: router,
	}

	router.POST("/user", controller.CreateNewUser)

	return &controller
}
