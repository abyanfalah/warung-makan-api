package model

type TransactionDetail struct {
	Id            string `json:"detail_id" db:"id"`
	TransactionId string `json:"transaction_id" db:"transaction_id"`
	ProductId     string `json:"product_id" binding:"required" db:"product_id"`
	Qty           int    `json:"qty" binding:"required" db:"qty"`
	Subtotal      int    `json:"subtotal" `
}
