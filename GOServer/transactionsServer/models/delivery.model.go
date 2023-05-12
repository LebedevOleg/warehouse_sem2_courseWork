package models

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
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Price float32 `json:"price"`
}

type DeliveryRequest struct {
	ProviderId int            `json:"provider"`
	StorageId  int            `json:"storage"`
	Items      []DeliveryItem `json:"items"`
}
