package database

import "database/sql"

func (db *DeliveryDB) GetAllProviders() (*sql.Rows, error) {
	rows, err := db.Db.Query(`SELECT id, name, address, phone FROM providers`)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
