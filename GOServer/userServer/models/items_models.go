package models

type Item struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Price float32 `json:"price"`
}
