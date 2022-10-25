package controller

import (
	"log"
	"net/http"
	"os"
	"warung-makan/config"
	"warung-makan/middleware"
	"warung-makan/model"
	"warung-makan/usecase"
	"warung-makan/utils"
	"warung-makan/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase usecase.UserUsecase
	router  *gin.Engine
}

func (c *UserController) ListUser(ctx *gin.Context) {
	if name := ctx.Query("name"); name != "" {
		user, err := c.usecase.GetByName(ctx.Query("name"))

		if err != nil {
			utils.JsonErrorNotFound(ctx, err, "cannot get list")
			return
		}

		if len(user) == 0 {
			ctx.String(http.StatusBadRequest, "no user with name like "+name)
			return
		}

		utils.JsonDataResponse(ctx, user)
		return
	}

	list, err := c.usecase.GetAll()
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "cannot get user list")
		return
	}

	utils.JsonDataResponse(ctx, list)
}

func (c *UserController) GetById(ctx *gin.Context) {
	user, err := c.usecase.GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorNotFound(ctx, err, "cannot get user")
		return
	}

	utils.JsonDataResponse(ctx, user)
}

func (c *UserController) CreateNewUser(ctx *gin.Context) {
	var user model.User
	c.router.MaxMultipartMemory = 8 << 20

	err := ctx.ShouldBind(&user)
	if err != nil {
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
		utils.JsonErrorInternalServerError(ctx, err, "cannot save image")
		return
	}

	user.Id = id
	user.Image = id + ".jpg"
	user, err = c.usecase.Insert(&user)
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "insert failed")
		return
	}

	utils.JsonDataMessageResponse(ctx, user, "user created")
}

func (c *UserController) CreateNewUserNoImage(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"data":  user,
		})
		return
	}
	// id := utils.GenerateId()

	// user.Id = id
	// user.Image = id + ".jpg"
	user, err = c.usecase.Insert(&user)
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "insert failed")
		return
	}

	// utils.JsonDataMessageResponse(ctx, &user, "user created")
	utils.JsonDataResponse(ctx, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	user.Id = ctx.Param("id")
	updatedUser, err := c.usecase.Update(&user)
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "update failed")
		return
	}

	utils.JsonDataResponse(ctx, updatedUser)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	user, err := c.usecase.GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorNotFound(ctx, err, "user not found")
		return
	}

	err = c.usecase.Delete(user.Id)
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "cannot delete user")
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
		utils.JsonErrorNotFound(ctx, err, imagePath)
		return
	}

	ctx.File(imagePath)
}

func NewUserController(usecase usecase.UserUsecase, router *gin.Engine) *UserController {
	controller := UserController{
		usecase: usecase,
		router:  router,
	}
	authMiddleware := middleware.NewAuthTokenMiddleware(authenticator.NewAccessToken(config.NewConfig().TokenConfig))

	router.GET("/user", controller.ListUser)
	router.GET("/user/:id", controller.GetById)
	router.GET("/user/:id/image", controller.GetUserImage)

	protectedRoute := router.Group("/user", authMiddleware.RequireToken())
	protectedRoute.POST("/", controller.CreateNewUser)
	protectedRoute.POST("/no_image", controller.CreateNewUserNoImage)
	protectedRoute.PUT("/:id", controller.UpdateUser)
	protectedRoute.DELETE("/:id", controller.DeleteUser)

	return &controller
}
