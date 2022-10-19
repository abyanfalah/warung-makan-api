package controller

import (
	"log"
	"net/http"
	"os"
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
	if name := ctx.Query("name"); name != "" {
		user, err := c.ucMan.UserUsecase().GetByName(ctx.Query("name"))

		if err != nil {
			utils.JsonErrorBadRequest(ctx, err, "cannot get list")
			return
		}

		if len(user) == 0 {
			ctx.String(http.StatusBadRequest, "no user with name like "+name)
			return
		}

		utils.JsonDataResponse(ctx, user)
		return
	}

	list, err := c.ucMan.UserUsecase().GetAll()
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "cannot get user list")
		return
	}

	utils.JsonDataResponse(ctx, list)
}

func (c *UserController) GetById(ctx *gin.Context) {
	user, err := c.ucMan.UserUsecase().GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cannot get user")
		return
	}

	utils.JsonDataResponse(ctx, user)
}

func (c *UserController) CreateNewUser(ctx *gin.Context) {
	var user model.User
	c.router.MaxMultipartMemory = 8 << 20

	err := ctx.ShouldBind(&user)
	if err != nil {
		// utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"data":  user,
		})
		return
	}

	imageFile, err := ctx.FormFile("image_file")
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant get image")
		return
	}

	id := utils.GenerateId()

	imagePath := "./images/user/" + id + ".jpg"
	err = ctx.SaveUploadedFile(imageFile, imagePath)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "cannot save image")
		return
	}

	user.Id = id
	user.Image = id + ".jpg"
	user, err = c.ucMan.UserUsecase().Insert(&user)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "insert failed")
		return
	}

	utils.JsonDataMessageResponse(ctx, user, "user created")
	ctx.JSON(200, gin.H{
		"uid": user.Id,
		"id":  id,
	})
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
		return
	}

	err = os.Remove("./images/user/" + user.Id + ".jpg")
	if err != nil {
		log.Println(err)
	}

	utils.JsonSuccessMessage(ctx, "User deleted")
}

func (c *UserController) GetUserImage(ctx *gin.Context) {
	id := ctx.Param("id")
	imagePath := "./images/user/" + id + ".jpg"
	if _, err := os.Stat(imagePath); err != nil {
		utils.JsonErrorBadRequest(ctx, err, imagePath)
		return
	}

	ctx.File(imagePath)
}

func NewUserController(usecaseManager manager.UsecaseManager, router *gin.Engine) *UserController {
	controller := UserController{
		ucMan:  usecaseManager,
		router: router,
	}
	authMiddleware := middleware.NewAuthTokenMiddleware(authenticator.NewAccessToken(config.NewConfig().TokenConfig))

	router.GET("/user", controller.ListUser)
	router.GET("/user/:id", controller.GetById)
	router.GET("/user/:id/image", controller.GetUserImage)

	protectedRoute := router.Group("/user", authMiddleware.RequireToken())
	protectedRoute.POST("/", controller.CreateNewUser)
	protectedRoute.PUT("/:id", controller.UpdateUser)
	protectedRoute.DELETE("/:id", controller.DeleteUser)

	return &controller
}
