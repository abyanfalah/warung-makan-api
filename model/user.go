package model

import "database/sql"

type User struct {
	Id       string         `json:"id" db:"id" `
	Name     string         `json:"name" db:"name" binding:"required"`
	Username string         `json:"username" db:"username" binding:"required"`
	Password string         `json:"password,omitempty" db:"password" binding:"required"`
	Image    sql.NullString `json:"image" db:"image"`

	// Image string `json:"image" db:"image" binding:"required"`
}

type Credential struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
