package models

import "database/sql"

type Delivery struct {
}

type Provider struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type Storage struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type DeliveryItem struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Count     int     `json:"count"`
	NeedCount int     `json:"need_count"`
	Price     float32 `json:"price"`
}

type DeliveryRequest struct {
	ProviderId int            `json:"provider"`
	StorageId  int            `json:"storage"`
	Items      []DeliveryItem `json:"items"`
}

type OrderJson struct {
	Id        int            `json:"id"`
	DateStart string         `json:"date_start"`
	DateEnd   sql.NullString `json:"date_end"`
	Status    int            `json:"status"`
	Price     float32        `json:"price"`
	Storage   Storage        `json:"storage"`
	User_id   int            `json:"user_id"`
}

type UserJwt struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	UType string `json:"user_type"`
}
