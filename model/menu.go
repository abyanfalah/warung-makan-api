package model

type Menu struct {
	Id    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name" binding:"required"`
	Price int    `json:"price" db:"price" binding:"required"`
	Stock int    `json:"stock" db:"stock" binding:"required"`
}
