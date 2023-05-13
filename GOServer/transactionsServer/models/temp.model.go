package models

type DeliveryTemp struct {
	Id             string
	ProviderName   string
	StorageAddress string
	Date           string
	Items          []ItemTemp
}

type ItemTemp struct {
	Id    string
	Name  string
	Count string
	Price string
}
