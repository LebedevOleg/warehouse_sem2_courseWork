package models

type DeliveryTemp struct {
	Id              int
	Provider_name   string
	Storage_address string
	Items           []DeliveryItem
}
