package model

type Menu struct {
	Id    string `json:"id" form:"id" db:"id"`
	Name  string `json:"name" form:"name" db:"name" binding:"required"`
	Price int    `json:"price" form:"price" db:"price" binding:"required"`
	Stock int    `json:"stock" form:"stock" db:"stock" binding:"required"`
	Image string `json:"image" form:"image" db:"image"`
}
