package controller

import (
	"net/http"

	"warung-makan/manager"
	"warung-makan/model"
	"warung-makan/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	ucMan  manager.UsecaseManager
	router *gin.Engine
}

func (c *UserController) CreateNewUser(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		utils.JsonErrorResponse(ctx, err, "cant bind struct")
		return
	}

	user.Id = utils.GenerateId()
	newUser, err := c.ucMan.UserUsecase().Insert(&user)
	if err != nil {
		utils.JsonErrorResponse(ctx, err, "insert failed")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":    newUser,
		"message": "new user created",
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
