package models

type DeliveryTemp struct {
	Id    string `json:"id"`
	Prov  string `json:"prov"`
	Adrs  string `json:"adrs"`
	Date  string `json:"date"`
	Text  TextTemp
	Items []ItemTemp `json:"items"`
}

type TextTemp struct {
	ProviderName   string
	Date           string
	StorageAddress string
}

type ItemTemp struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count string `json:"count"`
	Price string `json:"price"`
}
