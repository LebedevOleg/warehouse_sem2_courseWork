package services

import (
	"errors"
	"practice2sem/transactionsServer/database"
	"practice2sem/transactionsServer/models"
)

func GetAllProviders() ([]models.Provider, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("failed to connect to postgresql: " + err.Error())
	}
	rows, err := db.GetAllProviders()
	if err != nil {
		return nil, errors.New("failed to get providers: " + err.Error())
	}
	providersArr := make([]models.Provider, 0, 256)
	i := 0
	for rows.Next() {
		providersArr = append(providersArr, models.Provider{})
		err = rows.Scan(
			&providersArr[i].Id, &providersArr[i].Name,
			&providersArr[i].Address, &providersArr[i].Phone)
		if err != nil {
			return nil, errors.New("failed to scan providers" + err.Error())
		}
		i++
	}
	return providersArr, nil
}

func GetAllStorages() ([]models.Storage, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("failed to connect to postgresql: " + err.Error())
	}
	rows, err := db.GetAllStorages()
	if err != nil {
		return nil, errors.New("failed to get storages: " + err.Error())
	}
	storagesArr := make([]models.Storage, 0, 256)
	i := 0
	for rows.Next() {
		storagesArr = append(storagesArr, models.Storage{})
		err = rows.Scan(
			&storagesArr[i].Id, &storagesArr[i].Name,
			&storagesArr[i].Address)
		if err != nil {
			return nil, errors.New("failed to scan storages" + err.Error())
		}
		i++
	}
	return storagesArr, nil
}

func CreateNewDelivery(deliveryRequest models.DeliveryRequest) (*models.DeliveryTemp, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("failed to connect to postgresql: " + err.Error())
	}
	deliveryData, err := db.CreateNewDelivery(deliveryRequest)
	if err != nil {
		return nil, errors.New("failed to create new delivery: " + err.Error())
	}
	return deliveryData, nil
}

func CreateNewProvider(providerRequest models.Provider) error {
	db, err := database.GetPostgresql()
	if err != nil {
		return errors.New("failed to connect to postgresql: " + err.Error())
	}
	err = db.CreateNewProvider(providerRequest)
	if err != nil {
		return errors.New("failed to create new provider: " + err.Error())
	}
	return nil
}

func GetAllOrders() ([]models.OrderJson, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("failed to connect to postgresql: " + err.Error())
	}
	orders, err := db.GetAllOrders()
	if err != nil {
		return nil, errors.New("failed to get orders: " + err.Error())
	}
	return orders, nil
}

func GetOrderDetails(id int) ([]models.DeliveryItem, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("failed to connect to postgresql: " + err.Error())
	}
	order, err := db.GetOrderDetails(id)
	if err != nil {
		return nil, errors.New("failed to get order: " + err.Error())
	}
	return order, nil
}

func GetUsersOrders(id int) ([]models.OrderJson, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("failed to connect to postgresql: " + err.Error())
	}
	orders, err := db.GetUserOrders(id)
	if err != nil {
		return nil, errors.New("failed to get orders: " + err.Error())
	}
	return orders, nil
}

func UpdateOrder(newOder models.OrderJson) error {
	db, err := database.GetPostgresql()
	if err != nil {
		return errors.New("failed to connect to postgresql: " + err.Error())
	}
	err = db.UpdateOrder(newOder)
	if err != nil {
		return errors.New("failed to update order: " + err.Error())
	}
	return nil

}

func DeleteOrder(id int) error {
	db, err := database.GetPostgresql()
	if err != nil {
		return errors.New("failed to connect to postgresql: " + err.Error())
	}
	err = db.DeleteOrder(id)
	if err != nil {
		return errors.New("failed to delete order: " + err.Error())
	}
	return nil
}
