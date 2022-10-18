package model

type User struct {
	Id       string `json:"id" form:"id" db:"id" `
	Name     string `json:"name" form:"name" db:"name" binding:"required"`
	Username string `json:"username" form:"username" db:"username" binding:"required"`
	Password string `json:"password,,omitempty" form:"password" db:"password" binding:"required"`
	Image    string `json:"image" db:"image" form:"image"`
}

type Credential struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
