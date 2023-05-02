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
