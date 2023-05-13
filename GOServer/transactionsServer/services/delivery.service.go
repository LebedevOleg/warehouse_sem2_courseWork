package services

import (
	"errors"
	"path/filepath"
	"practice2sem/transactionsServer/database"
	"practice2sem/transactionsServer/models"
	"time"

	docxt "github.com/legion-zver/go-docx-templates"
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

func CreateNewDelivery(deliveryRequest models.DeliveryRequest) (string, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return "", errors.New("failed to connect to postgresql: " + err.Error())
	}
	deliveryData, err := db.CreateNewDelivery(deliveryRequest)
	if err != nil {
		return "", errors.New("failed to create new delivery: " + err.Error())
	}
	template, err := docxt.OpenTemplate(filepath.Join(".", "templates", "Delivery_temp.docx"))
	if err != nil {
		return "", errors.New("failed to open template: " + err.Error())
	}

	err = template.RenderTemplate(&deliveryData)
	if err != nil {
		return "", errors.New("failed to render template: " + err.Error())
	}
	fileName := "delivery_" + time.Now().Format("2006-01-02_15-04-05") + ".docx"
	err = template.Save(filepath.Join(".", "files", fileName))
	if err != nil {
		return "", errors.New("failed to save file: " + err.Error())
	}
	return fileName, nil
}
