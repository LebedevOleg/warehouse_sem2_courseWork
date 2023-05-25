package models

type Item struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Price float32 `json:"price"`
}
type Offer struct {
	StorageId int    `json:"storage_id"`
	Items     []Item `json:"items"`
}
