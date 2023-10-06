package models

type User struct {
	User_id    string `db:"user_id" form:"user_id" valid:"-" json:"user_id,omitempty"`
	Name       string `db:"name" form:"name" json:"name" valid:"-" `
	Created_at string `db:"created_at" json:"created_at" `
}
