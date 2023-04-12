package services

import (
	"errors"
	"practice2sem/itemsServer/database"
	"practice2sem/itemsServer/models"
)

func CreateItem(item *models.ItemJson) error {
	db, err := database.GetPostgresql()
	if err != nil {
		return errors.New("failed to connect to postgresql: " + err.Error())
	}
	err = db.CreateItem(*item)
	if err != nil {
		return errors.New("failed to create item: " + err.Error())
	}
	return nil
}

func GetItem(id int) (*models.ItemJson, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("failed to connect to postgresql: " + err.Error())
	}
	row := db.GetItem(id)
	var resultItem models.ItemJson
	err = row.Scan(
		&resultItem.Id, &resultItem.Name,
		&resultItem.Descriptions, &resultItem.Category_id,
		&resultItem.Category_name, &resultItem.Price_for_unit,
		&resultItem.Dimension)
	if err != nil {
		return nil, errors.New("failed to get item: " + err.Error())
	}
	return &resultItem, nil
}

func GetAllItems() ([]models.ItemJson, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("failed to connect to postgresql: " + err.Error())
	}
	rows, err := db.GetAllItems()
	if err != nil {
		return nil, errors.New("failed to get items: " + err.Error())
	}
	itemsArr := make([]models.ItemJson, 0, 256)
	i := 0
	for rows.Next() {
		itemsArr = append(itemsArr, models.ItemJson{})
		err = rows.Scan(
			&itemsArr[i].Id, &itemsArr[i].Name,
			&itemsArr[i].Descriptions, &itemsArr[i].Category_id,
			&itemsArr[i].Category_name, &itemsArr[i].Price_for_unit,
			&itemsArr[i].Dimension)
		if err != nil {
			return nil, errors.New("failed to scan items" + err.Error())
		}
		i++
	}
	return itemsArr, nil
}
