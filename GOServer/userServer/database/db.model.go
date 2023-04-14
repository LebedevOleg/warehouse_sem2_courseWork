package database

import (
	"database/sql"
	"practice2sem/server/database"
	"practice2sem/userServer/models"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0000"
	dbName   = "warehouse"
)

func GetPostgresql() (*UserDB, error) {
	if postgresqlDB == nil {
		postgresqlDB = &UserDB{}
		postgresqlDB.SetParam(dbName, host, password, user, port)
		err := postgresqlDB.Connect()
		if err != nil {
			return postgresqlDB, err
		}
		return postgresqlDB, nil
	}
	return postgresqlDB, nil
}

var postgresqlDB *UserDB

type UserDB struct {
	database.Postgresql
}

func (p *UserDB) CreateUser(u models.UserJson) error {
	_, err := p.Db.Exec("INSERT into users (name, password, email) VALUES ($1, $2, $3)",
		u.Name, u.Password, u.Email)
	if err != nil {
		return err
	}
	return nil
}

func (p *UserDB) GetUser(u models.UserJson) *sql.Row {
	row := p.Db.QueryRow(`SELECT users.id, email, password, t.name FROM users, user_types t where type_id = t.id and email = $1`, u.Email)
	return row
}
