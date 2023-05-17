package models

type DeliveryTemp struct {
	Id    string
	Prov  string
	Adrs  string
	Date  string
	Text  TextTemp
	Items []ItemTemp
}

type TextTemp struct {
	ProviderName   string
	Date           string
	StorageAddress string
}

type ItemTemp struct {
	Id    string
	Name  string
	Count string
	Price string
}
