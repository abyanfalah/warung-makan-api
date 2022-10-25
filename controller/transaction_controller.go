package controller

import (
	"fmt"
	"net/http"
	"warung-makan/config"
	"warung-makan/middleware"
	"warung-makan/model"
	"warung-makan/usecase"
	"warung-makan/utils"
	"warung-makan/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	usecase     usecase.TransactionUsecase
	menuUsecase usecase.MenuUsecase
	router      *gin.Engine
}

func (c *TransactionController) ListTransaction(ctx *gin.Context) {
	list, err := c.usecase.GetAll()
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "cannot get transaction list")
		return
	}

	utils.JsonDataResponse(ctx, list)
}

func (c *TransactionController) GetById(ctx *gin.Context) {
	transaction, err := c.usecase.GetById(ctx.Param("id"))
	if err != nil {
		utils.JsonErrorBadRequest(ctx, err, "cannot get transaction")
		return
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
		menu, err := c.menuUsecase.GetById(each.MenuId)
		if err != nil {
			fmt.Println("cant find the menu")
			continue
		}

		if each.Qty > menu.Stock || menu.Stock < 1 || each.Qty < 1 {
			fmt.Println("qty is invalid")
			continue
		}
		transaction.Items[i].TransactionId = transaction.Id
		transaction.Items[i].Subtotal = menu.Price * each.Qty
		transaction.TotalPrice += transaction.Items[i].Subtotal
	}

	if len(transaction.Items) == 0 {
		ctx.String(http.StatusBadRequest, "Transaction has 0 valid item. Transaction not created.")
		return
	}

	newTransaction, err := c.usecase.Insert(&transaction)
	if err != nil {
		utils.JsonErrorInternalServerError(ctx, err, "insert failed")
		return
	}

	utils.JsonDataMessageResponse(ctx, newTransaction, "transaction created")
}

// func (c *TransactionController) UpdateTransaction(ctx *gin.Context) {
// 	var transaction model.Transaction

// 	err := ctx.ShouldBindJSON(&transaction)
// 	if err != nil {
// 		utils.JsonErrorBadRequest(ctx, err, "cant bind struct")
// 		return
// 	}

// 	transaction.Id = ctx.Param("id")
// 	updatedTransaction, err := c.usecase.Update(&transaction)
// 	if err != nil {
// 		utils.JsonErrorInternalServerError(ctx, err, "update failed")
// 		return
// 	}

// 	utils.JsonDataMessageResponse(ctx, updatedTransaction, "transaction updated")
// }

// func (c *TransactionController) DeleteTransaction(ctx *gin.Context) {
// 	transaction, err := c.usecase.GetById(ctx.Param("id"))
// 	if err != nil {
// 		utils.JsonErrorBadRequest(ctx, err, "transaction not found")
// 		return
// 	}

// 	err = c.usecase.Delete(transaction.Id)
// 	if err != nil {
// 		utils.JsonErrorInternalServerError(ctx, err, "cannot delete transaction")
// 	}

// 	utils.JsonSuccessMessage(ctx, "Transaction deleted")
// }

func NewTransactionController(usecase usecase.TransactionUsecase, menuUsecase usecase.MenuUsecase, router *gin.Engine) *TransactionController {
	controller := TransactionController{
		usecase:     usecase,
		menuUsecase: menuUsecase,
		router:      router,
	}
	authMiddleware := middleware.NewAuthTokenMiddleware(authenticator.NewAccessToken(config.NewConfig().TokenConfig))

	protectedRoute := router.Group("/transaction", authMiddleware.RequireToken())
	protectedRoute.GET("", controller.ListTransaction)
	protectedRoute.GET("/:id", controller.GetById)
	protectedRoute.POST("", controller.CreateNewTransaction)

	return &controller
}
