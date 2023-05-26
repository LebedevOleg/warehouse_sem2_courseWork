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

func (db *DeliveryDB) GetAllOrders() ([]models.OrderJson, error) {
	order := make([]models.OrderJson, 0, 254)
	rows, err := db.Db.Query(`SELECT id, date_start, date_finish, status, price, storage_id FROM orders`)
	if err != nil {
		return nil, err
	}
	i := 0
	for rows.Next() {
		order = append(order, models.OrderJson{})
		err = rows.Scan(
			&order[i].Id, &order[i].DateStart, &order[i].DateEnd, &order[i].Status, &order[i].Price, &order[i].Storage.Id,
		)
		if err != nil {
			return nil, err
		}
		err = db.Db.QueryRow(`SELECT name, address FROM storages WHERE id = $1`,
			&order[i].Storage.Id).Scan(&order[i].Storage.Name, &order[i].Storage.Address)
		if err != nil {
			return nil, err
		}
		i++
	}
	return order, nil
}

func (db *DeliveryDB) GetOrderDetails(id int) ([]models.DeliveryItem, error) {
	rows, err := db.Db.Query(`SELECT item_id, item_count, item_need_count, i.name from items_to_orders ito, items i 
		WHERE ito.order_id = $1`, id)
	if err != nil {
		return nil, err
	}
	items := make([]models.DeliveryItem, 0, 254)
	i := 0
	for rows.Next() {
		items = append(items, models.DeliveryItem{})
		err = rows.Scan(
			&items[i].Id, &items[i].Count, &items[i].NeedCount, &items[i].Name,
		)
		if err != nil {
			return nil, err
		}
		i++
	}
	return items, nil
}

func (db *DeliveryDB) GetUserOrders(userId int) ([]models.OrderJson, error) {
	var orders []models.OrderJson
	rows, err := db.Db.Query(`SELECT id, date_start, date_finish, status, price, storage_id FROM orders WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var order models.OrderJson
		err = rows.Scan(
			&order.Id, &order.DateStart, &order.DateEnd, &order.Status, &order.Price, &order.Storage.Id,
		)
		if err != nil {
			return nil, err
		}
		err = db.Db.QueryRow(`SELECT name, address FROM storages WHERE id = $1`,
			&order.Storage.Id).Scan(&order.Storage.Name, &order.Storage.Address)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (db *DeliveryDB) UpdateOrder(order models.OrderJson) error {
	_, err := db.Db.Exec(
		`UPDATE orders SET date_start = $1, date_finish = $2, status = $3, price = $4 WHERE id = $5`,
		order.DateStart, order.DateEnd, order.Status, order.Price, order.Id)
	if err != nil {
		return err
	}
	return nil
}

func (db *DeliveryDB) DeleteOrder(id int) error {
	_, err := db.Db.Exec(`DELETE FROM orders WHERE id = $1`, id)
	if err != nil {
		return err
	}
	_, err = db.Db.Exec(`DELETE FROM providers_to_deliveries WHERE order_id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
