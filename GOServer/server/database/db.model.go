package database

import (
	"database/sql"
	"fmt"
)

type Postgresql struct {
	dbName     string
	dbHost     string
	dbPort     int
	dbPassword string
	dbUser     string
	Db         *sql.DB
}

func (p *Postgresql) SetParam(name, host, pass, user string, port int) {
	p.dbHost = host
	p.dbName = name
	p.dbPassword = pass
	p.dbPort = port
	p.dbUser = user
}

func (p *Postgresql) Connect() error {
	conf := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.dbHost, p.dbPort, p.dbUser, p.dbPassword, p.dbName)
	db, err := sql.Open("postgres", conf)
	if err != nil {
		return err
	}
	p.Db = db
	err = p.Db.Ping()
	if err != nil {
		return err
	}
	return nil
}
