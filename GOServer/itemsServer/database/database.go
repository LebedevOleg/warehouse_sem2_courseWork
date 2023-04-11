package database

import (
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

//id	name	description	category_id	price_for_unit	dimension
