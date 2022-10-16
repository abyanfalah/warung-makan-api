package model

type TransactionDetail struct {
	TransactionId string `json:"transaction_id" db:"transaction_id"`
	MenuId        string `json:"menu_id" binding:"required" db:"menu_id"`
	Qty           int    `json:"qty" binding:"required" db:"qty"`
	Subtotal      int    `json:"subtotal" `
}
