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

type MenuController struct {
	ucMan  manager.UsecaseManager
	router *gin.Engine
}

func (c *MenuController) ListMenu(ctx *gin.Context) {
	list, err := c.ucMan.MenuUsecase().GetAll()
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "cannot get menu list")
	}

	utils.JsonDataResponse(ctx, list)
}

func (c *MenuController) GetById(ctx *gin.Context) {
	menu, err := c.ucMan.MenuUsecase().GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cannot get menu")
	}

	utils.JsonDataResponse(ctx, menu)
}

func (c *MenuController) CreateNewMenu(ctx *gin.Context) {
	var menu model.Menu
	c.router.MaxMultipartMemory = 8 << 20

	err := ctx.ShouldBind(&menu)
	if err != nil {
		// utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		utils.JsonDataMessageResponse(ctx, menu, "cant bind struct: "+err.Error())
		return
	}

	imageFile, err := ctx.FormFile("image_file")
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant get image")
	}

	// utils.JsonDataResponse(ctx, gin.H{
	// 	"menu":  menu,
	// 	"image": imageFile,
	// })
	// return

	menu.Id = utils.GenerateId()
	imagePath := "./images/menu/" + menu.Id + ".jpg"
	err = ctx.SaveUploadedFile(imageFile, imagePath)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "cannot save image")
		return
	}

	menu.Image = imagePath
	newMenu, err := c.ucMan.MenuUsecase().Insert(&menu)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "insert failed")
		return
	}

	utils.JsonDataMessageResponse(ctx, newMenu, "menu created")
}

func (c *MenuController) UpdateMenu(ctx *gin.Context) {
	var menu model.Menu

	err := ctx.ShouldBindJSON(&menu)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	menu.Id = ctx.Param("id")
	updatedMenu, err := c.ucMan.MenuUsecase().Update(&menu)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "update failed")
		return
	}

	utils.JsonDataMessageResponse(ctx, updatedMenu, "menu updated")
}

func (c *MenuController) DeleteMenu(ctx *gin.Context) {
	menu, err := c.ucMan.MenuUsecase().GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "menu not found")
		return
	}

	err = c.ucMan.MenuUsecase().Delete(menu.Id)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "cannot delete menu")
	}

	utils.JsonSuccessMessage(ctx, "Menu deleted")
}

func NewMenuController(usecaseManager manager.UsecaseManager, router *gin.Engine) *MenuController {
	controller := MenuController{
		ucMan:  usecaseManager,
		router: router,
	}
	authMiddleware := middleware.NewAuthTokenMiddleware(authenticator.NewAccessToken(config.NewConfig().TokenConfig))

	router.GET("/menu", controller.ListMenu)
	router.GET("/menu/:id", controller.GetById)

	protectedRoute := router.Group("/menu", authMiddleware.RequireToken())
	protectedRoute.POST("/", controller.CreateNewMenu)
	protectedRoute.PUT("/:id", controller.UpdateMenu)
	protectedRoute.DELETE("/:id", controller.DeleteMenu)

	return &controller
}
