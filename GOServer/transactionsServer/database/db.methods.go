package database

import (
	"database/sql"
	"fmt"
	"practice2sem/transactionsServer/models"
	"strconv"
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
		`INSERT INTO deliveries (storage_id, delivery_date) VALUES($1, $2) RETURNING id, delivery_date`,
		deliveryRequest.StorageId, time.Now()).Scan(
		&deliveryTemp.Id,
		&deliveryTemp.Date,
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
	items := make([]models.ItemTemp, 0, len(deliveryRequest.Items))
	for _, item := range deliveryRequest.Items {
		err = db.Db.QueryRow(
			`SELECT price_for_unit FROM items WHERE name = $1`,
			item.Name).Scan(
			&item.Price,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, models.ItemTemp{
			Id:    strconv.Itoa(item.Id),
			Name:  item.Name,
			Count: strconv.Itoa(item.Count),
			Price: fmt.Sprintf("%.2f", item.Price*float32(item.Count)),
		})
	}
	deliveryTemp.Items = items
	err = db.Db.QueryRow(
		`SELECT name FROM providers WHERE id = $1`,
		deliveryRequest.ProviderId).Scan(
		&deliveryTemp.Prov,
	)
	if err != nil {
		return nil, err
	}
	err = db.Db.QueryRow(
		`SELECT address FROM storages WHERE id = $1`,
		deliveryRequest.StorageId).Scan(
		&deliveryTemp.Adrs,
	)
	if err != nil {
		return nil, err
	}
	deliveryTemp.Text = models.TextTemp{
		ProviderName:   deliveryTemp.Prov,
		Date:           deliveryTemp.Date,
		StorageAddress: deliveryTemp.Adrs,
	}
	return deliveryTemp, nil
}

func (db *DeliveryDB) CreateNewProvider(provider models.Provider) error {
	_, err := db.Db.Exec(
		`INSERT INTO providers (name, address, phone) VALUES($1, $2, $3)`,
		provider.Name, provider.Address, provider.Phone)
	if err != nil {
		return err
	}
	return nil
}
