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

func (c *UserController) ListUser(ctx *gin.Context) {
	list, err := c.ucMan.UserUsecase().GetAll()
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "cannot get user list")
	}

	utils.JsonDataResponse(ctx, list)
}

func (c *UserController) GetById(ctx *gin.Context) {
	user, err := c.ucMan.UserUsecase().GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cannot get user")
	}

	utils.JsonDataResponse(ctx, user)
}

func (c *UserController) CreateNewUser(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	user.Id = utils.GenerateId()
	newUser, err := c.ucMan.UserUsecase().Insert(&user)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "insert failed")
		return
	}

	utils.JsonDataResponse(ctx, newUser)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	user.Id = ctx.Param("id")
	updatedUser, err := c.ucMan.UserUsecase().Update(&user)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "insert failed")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"updated_user": updatedUser,
		"message":      "user updated",
	})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	user, err := c.ucMan.UserUsecase().GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "user not found")
		return
	}

	err = c.ucMan.UserUsecase().Delete(user.Id)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "cannot delete user")
	}

	utils.JsonSuccessMessage(ctx, "User deleted")
}

func NewUserController(usecaseManager manager.UsecaseManager, router *gin.Engine) *UserController {
	controller := UserController{
		ucMan:  usecaseManager,
		router: router,
	}

	router.GET("/user", controller.ListUser)
	router.GET("/user/:id", controller.GetById)
	router.POST("/user", controller.CreateNewUser)
	router.PUT("/user/:id", controller.UpdateUser)
	router.DELETE("/user/:id", controller.DeleteUser)

	return &controller
}
