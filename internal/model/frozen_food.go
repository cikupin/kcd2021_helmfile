package model

type FrozenFood struct {
	ID       int    `json:"id" db:"id,primarykey"`
	Name     string `json:"name" db:"name"`
	Quantity int    `json:"quantity" db:"quantity"`
}
