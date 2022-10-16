package controller

import (
	"warung-makan/manager"
	"warung-makan/model"
	"warung-makan/utils"

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

	err := ctx.ShouldBindJSON(&menu)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	menu.Id = utils.GenerateId()
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

	router.GET("/menu", controller.ListMenu)
	router.GET("/menu/:id", controller.GetById)
	router.POST("/menu", controller.CreateNewMenu)
	router.PUT("/menu/:id", controller.UpdateMenu)
	router.DELETE("/menu/:id", controller.DeleteMenu)

	return &controller
}
