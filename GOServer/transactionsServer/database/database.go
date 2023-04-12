package database

import "practice2sem/server/database"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0000"
	dbName   = "warehouse"
)

func GetPostgresql() (*DeliveryDB, error) {
	if postgresqlDB == nil {
		postgresqlDB = &DeliveryDB{}
		postgresqlDB.SetParam(dbName, host, password, user, port)
		err := postgresqlDB.Connect()
		if err != nil {
			return postgresqlDB, err
		}
		return postgresqlDB, nil
	}
	return postgresqlDB, nil
}

var postgresqlDB *DeliveryDB

type DeliveryDB struct {
	database.Postgresql
}
