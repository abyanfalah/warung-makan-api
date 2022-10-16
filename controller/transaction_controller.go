package controller

import (
	"warung-makan/manager"
	"warung-makan/model"
	"warung-makan/utils"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	ucMan  manager.UsecaseManager
	router *gin.Engine
}

func (c *TransactionController) ListTransaction(ctx *gin.Context) {
	list, err := c.ucMan.TransactionUsecase().GetAll()
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "cannot get transaction list")
	}

	utils.JsonDataResponse(ctx, list)
}

func (c *TransactionController) GetById(ctx *gin.Context) {
	transaction, err := c.ucMan.TransactionUsecase().GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cannot get transaction")
	}

	utils.JsonDataResponse(ctx, transaction)
}

func (c *TransactionController) CreateNewTransaction(ctx *gin.Context) {
	var transaction model.Transaction

	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	transaction.Id = utils.GenerateId()
	// fill transaction details
	for i, each := range transaction.Items {
		menu, _ := c.ucMan.MenuUsecase().GetById(each.MenuId)

		if each.Qty > menu.Stock || menu.Stock < 1 || each.Qty < 1 {
			continue
		}
		transaction.Items[i].TransactionId = transaction.Id
		transaction.Items[i].Subtotal = menu.Price * each.Qty
		transaction.TotalPrice += transaction.Items[i].Subtotal
	}

	newTransaction, err := c.ucMan.TransactionUsecase().Insert(&transaction)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "insert failed")
		return
	}

	utils.JsonDataMessageResponse(ctx, newTransaction, "transaction created")
}

func (c *TransactionController) UpdateTransaction(ctx *gin.Context) {
	var transaction model.Transaction

	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
		return
	}

	transaction.Id = ctx.Param("id")
	updatedTransaction, err := c.ucMan.TransactionUsecase().Update(&transaction)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "update failed")
		return
	}

	utils.JsonDataMessageResponse(ctx, updatedTransaction, "transaction updated")
}

func (c *TransactionController) DeleteTransaction(ctx *gin.Context) {
	transaction, err := c.ucMan.TransactionUsecase().GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "transaction not found")
		return
	}

	err = c.ucMan.TransactionUsecase().Delete(transaction.Id)
	if err != nil {
		utils.JsonErrorBadGateway(ctx, err, "cannot delete transaction")
	}

	utils.JsonSuccessMessage(ctx, "Transaction deleted")
}

func NewTransactionController(usecaseManager manager.UsecaseManager, router *gin.Engine) *TransactionController {
	controller := TransactionController{
		ucMan:  usecaseManager,
		router: router,
	}

	router.GET("/transaction", controller.ListTransaction)
	router.GET("/transaction/:id", controller.GetById)
	router.POST("/transaction", controller.CreateNewTransaction)
	router.PUT("/transaction/:id", controller.UpdateTransaction)
	router.DELETE("/transaction/:id", controller.DeleteTransaction)

	return &controller
}
