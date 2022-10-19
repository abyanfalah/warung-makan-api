package model

import "database/sql"

type Transaction struct {
	Id         string              `json:"id"`
	TotalPrice int                 `json:"total" db:"total_price"`
	Created_at string              `db:"created_at" json:"created_at"`
	Updated_at sql.NullTime        `db:"updated_at" json:"updated_at,omitempty"`
	Items      []TransactionDetail `json:"items" binding:"required"`
}
