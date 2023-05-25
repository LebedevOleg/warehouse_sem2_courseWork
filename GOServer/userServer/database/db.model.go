package database

import (
	"database/sql"
	"errors"
	"fmt"
	"practice2sem/server/database"
	"practice2sem/userServer/models"
	"time"

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

func (p *UserDB) CreateUser(u models.UserJson) (*sql.Row, error) {
	_, err := p.Db.Exec("INSERT into users (name, password, email) VALUES ($1, $2, $3)",
		u.Name, u.Password, u.Email)
	if err != nil {
		return nil, err
	}
	return p.GetUser(u), nil
}

// todo: check
func (p *UserDB) UpdateUser(u models.UserJson) error {
	_, err := p.Db.Exec(`UPDATE users SET name = $1, email = $2, password =$3, role = $4`, u.Name, u.Email, u.Password, u.Role)
	if err != nil {
		return errors.New("Error updating user " + err.Error())
	}
	return nil
}

func (p *UserDB) GetUser(u models.UserJson) *sql.Row {
	row := p.Db.QueryRow(`SELECT users.id, email, password, r.name FROM users, user_roles r where role_id = r.id and email = $1`, u.Email)
	return row
}

func (p *UserDB) GetAllUsers() (*sql.Rows, error) {
	rows, err := p.Db.Query(`SELECT users.id, users.name, email, r.name, t.name 
	FROM users, user_roles r, user_types t
	WHERE role_id = r.id and type_id = t.id`)
	if err != nil {
		return nil, err
	}
	return rows, nil

}

func (p *UserDB) GetUserById(id int) *sql.Row {
	row := p.Db.QueryRow(`SELECT * FROM users WHERE id = $1`, id)
	return row
}

func (p *UserDB) CreateOffer(u models.UserJwt, offer models.Offer) error {
	var allPrice float32
	allPrice = 0
	for _, item := range offer.Items {
		allPrice += item.Price * float32(item.Count)
	}
	id := 0
	err := p.Db.QueryRow(`INSERT INTO orders (date_start,  status, user_id, price, storage_id)
		VALUES($1, $2, $3, $4, $5) RETURNING id`,
		time.Now(), 0, u.Id, allPrice, offer.StorageId).Scan(&id)

	if err != nil {
		return errors.New("Error creating offer " + err.Error())
	}
	fmt.Println(id)
	for _, item := range offer.Items {
		_, err := p.Db.Exec(`INSERT INTO items_to_orders (order_id, item_id, item_count) VALUES($1, $2, $3)`,
			id, item.Id, item.Count)
		if err != nil {
			return errors.New("Ошибка в добавлении товара в промежуточную таблицу " + err.Error())
		}
	}
	return nil
}
