package controller

import (
	"warung-makan/config"
	"warung-makan/manager"
	"warung-makan/middleware"
	"warung-makan/model"
	"warung-makan/utils"
	"warung-makan/utils/authenticator"

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

	// duplicate, err := c.ucMan.UserUsecase().GetById(user.Id)
	// if err != nil {
	// 	utils.JsonErrorBadRequest(ctx, err, "user duplicate")
	// }

	user.Id = utils.GenerateId()
	newUser, err := c.ucMan.UserUsecase().Insert(&user)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "insert failed")
		return
	}

	utils.JsonDataMessageResponse(ctx, newUser, "user created")
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
		utils.JsonErrorBadGateway(ctx, err, "update failed")
		return
	}

	utils.JsonDataMessageResponse(ctx, updatedUser, "user updated")
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
	authMiddleware := middleware.NewAuthTokenMiddleware(authenticator.NewAccessToken(config.NewConfig().TokenConfig))

	router.GET("/user", controller.ListUser)
	router.GET("/user/:id", controller.GetById)

	protectedRoute := router.Group("/user", authMiddleware.RequireToken())
	protectedRoute.POST("/", controller.CreateNewUser)
	protectedRoute.PUT("/:id", controller.UpdateUser)
	protectedRoute.DELETE("/:id", controller.DeleteUser)

	return &controller
}
