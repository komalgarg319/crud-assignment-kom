package model

type Book struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	AuthID uint   `json:"auth_id"`
}
