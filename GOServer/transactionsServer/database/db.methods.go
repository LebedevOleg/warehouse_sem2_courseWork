package database

import (
	"database/sql"
	"practice2sem/transactionsServer/models"
	"time"
)

func (db *DeliveryDB) GetAllProviders() (*sql.Rows, error) {
	rows, err := db.Db.Query(`SELECT id, name, address, phone FROM providers`)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *DeliveryDB) GetAllStorages() (*sql.Rows, error) {
	rows, err := db.Db.Query(`SELECT id, name, address FROM storages`)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *DeliveryDB) CreateNewDelivery(deliveryRequest models.DeliveryRequest) (*models.DeliveryTemp, error) {
	var deliveryTemp = new(models.DeliveryTemp)
	err := db.Db.QueryRow(
		`INSERT INTO deliveries (storage_id, delivery_date) VALUES($1, $2) RETURNING id`,
		deliveryRequest.StorageId, time.Now().).Scan(
		deliveryTemp.Id,
	)
	if err != nil {
		return nil, err
	}
	for _, item := range deliveryRequest.Items {
		_, err = db.Db.Exec(`INSERT INTO providers_to_deliveries (provider_id, delivery_id, item_id, item_count) VALUES($1, $2, $3, $4)`,
			deliveryRequest.ProviderId, deliveryTemp.Id, item.Id, item.Count)
		if err != nil {
			return nil, err
		}
	}
	deliveryTemp.Items = deliveryRequest.Items
	err = db.Db.QueryRow(
		`SELECT name FROM providers WHERE id = $1`,
		deliveryRequest.ProviderId).Scan(
		&deliveryTemp.Provider_name,
	)
	if err != nil {
		return nil, err
	}
	err = db.Db.QueryRow(
		`SELECT address FROM storages WHERE id = $1`,
		deliveryRequest.StorageId).Scan(
		&deliveryTemp.Storage_address,
	)
	if err != nil {
		return nil, err
	}
	return deliveryTemp, nil
}
