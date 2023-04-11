package database

import (
	"database/sql"
	"fmt"
	"practice2sem/itemsServer/models"
	"practice2sem/server/database"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0000"
	dbName   = "warehouse"
)

func GetPostgresql() (*ItemDB, error) {
	if postgresqlDB == nil {
		postgresqlDB = &ItemDB{}
		postgresqlDB.SetParam(dbName, host, password, user, port)
		err := postgresqlDB.Connect()
		if err != nil {
			return postgresqlDB, err
		}
		return postgresqlDB, nil
	}
	return postgresqlDB, nil
}

var postgresqlDB *ItemDB

type ItemDB struct {
	database.Postgresql
}

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
	fmt.Println("\n", id)
	row := db.Db.QueryRow(`SELECT i.id, i.name, description, i.category_id, c.name, price_for_unit, dimension
		FROM items i, item_categories c
		WHERE i.category_id = c.id and i.id = $1`, id)
	return row
}

//id	name	description	category_id	price_for_unit	dimension
