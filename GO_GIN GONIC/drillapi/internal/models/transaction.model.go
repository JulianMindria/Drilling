package models

type Transaction struct {
	Transaction_id string `db:"transaction_id" form:"transaction_id" valid:"-" json:"transaction_id,omitempty"`
	Name           string `db:"name" form:"name" valid:"-" json:"name,omitempty"`
	Total          string `db:"total" form:"total" json:"total,omitempty"`
	User
}
