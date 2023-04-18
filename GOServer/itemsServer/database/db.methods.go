package database

import (
	"database/sql"
	"practice2sem/itemsServer/models"
)

func (db *ItemDB) CreateItem(i models.ItemJson) error {
	_, err := db.Db.Exec(`INSERT 
		INTO items (name, description, category_id, price_for_unit, dimension) 
		VALUES ($1, $2,$3,$4,$5)`,
		i.Name, i.Descriptions, i.Category_id, i.Price_for_unit, i.Dimension)
	if err != nil {
		return err
	}
	return nil
}

func (db *ItemDB) GetItem(id int) *sql.Row {
	row := db.Db.QueryRow(`SELECT i.id, i.name, description, i.category_id, c.name, price_for_unit, dimension
		FROM items i, item_categories c
		WHERE i.category_id = c.id and i.id = $1`, id)
	return row
}

func (db *ItemDB) GetAllItems() (*sql.Rows, error) {
	rows, err := db.Db.Query(`SELECT i.id, i.name, description, i.category_id, c.name, price_for_unit, dimension
        FROM items i, item_categories c
        WHERE i.category_id = c.id`)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *ItemDB) GetAllCategories() (*sql.Rows, error) {
	rows, err := db.Db.Query(`SELECT id, name
        FROM item_categories`)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *ItemDB) UpdateItem(i models.ItemJson) error {
	_, err := db.Db.Exec(`UPDATE items SET description = $1, price_for_unit = $2, dimension = $3, name = $4 WHERE id = $5`,
		i.Descriptions, i.Price_for_unit, i.Dimension, i.Name, i.Id)
	if err != nil {
		return err
	}
	return nil
}
