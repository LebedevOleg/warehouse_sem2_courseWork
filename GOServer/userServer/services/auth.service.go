package services

import (
	"errors"
	"practice2sem/userServer/database"
	"practice2sem/userServer/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Register(user *models.UserJson) (string, string, error) {
	err := user.HashPassword()
	if err != nil {
		return "", "", errors.New("hash password error\n" + err.Error())
	}
	db, err := database.GetPostgresql()
	if err != nil {
		return "", "", errors.New("get postgresql error\n" + err.Error())
	}

	row, err := db.CreateUser(*user)
	if err != nil {
		return "", "", errors.New("create user error\n" + err.Error())
	}
	var password string
	var email string
	var uType string
	var id int
	err = row.Scan(&id, &email, &password, &uType)
	if err != nil {
		return "", "", errors.New("get user error\n" + err.Error())
	}
	claims := &models.UserJwt{
		id,
		email,
		uType,
		jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 4))},
	}
	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := userToken.SignedString([]byte("secret"))
	if err != nil {
		return "", "", errors.New("sign token error\n" + err.Error())
	}
	return t, uType, nil
}

func Login(user *models.UserJson) (string, string, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return "", "", errors.New("get postgresql error\n" + err.Error())
	}
	row := db.GetUser(*user)
	var password string
	var email string
	var uType string
	var id int
	err = row.Scan(&id, &email, &password, &uType)
	if err != nil {
		return "", "", errors.New("get user error\n" + err.Error())
	}
	if !user.CheckPasswordHash(password) {
		return "", "", errors.New("password unmatch\n")
	}
	claims := &models.UserJwt{
		id,
		email,
		uType,
		jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 4))},
	}
	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := userToken.SignedString([]byte("secret"))
	if err != nil {
		return "", "", errors.New("sign token error\n" + err.Error())
	}
	return t, uType, nil
}

func GetAllUsers() ([]*models.UserJson, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("get postgresql error\n" + err.Error())
	}
	rows, err := db.GetAllUsers()
	if err != nil {
		return nil, errors.New("get all users error\n" + err.Error())
	}
	var users []*models.UserJson
	for rows.Next() {
		user := new(models.UserJson)
		err := rows.Scan(&user.Id, &user.Email, &user.Type, &user.Role, &user.Name)
		if err != nil {
			return nil, errors.New("get all users error\n" + err.Error())
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserInfo(id int) (*models.UserJson, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("get postgresql error\n" + err.Error())
	}
	row := db.GetUserById(id)
	user := new(models.UserJson)
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Role)
	if err != nil {
		return nil, errors.New("get user error\n" + err.Error())
	}
	return user, nil
}

func UpdateUserInfo(user *models.UserJson) error {
	db, err := database.GetPostgresql()
	if err != nil {
		return errors.New("get postgresql error\n" + err.Error())
	}
	err = db.UpdateUser(*user)
	if err != nil {
		return errors.New("update user error\n" + err.Error())
	}
	return nil
}

func CreateOffer(user *models.UserJwt, offer *models.Offer) error {
	db, err := database.GetPostgresql()
	if err != nil {
		return errors.New("get postgresql error\n" + err.Error())
	}
	err = db.CreateOffer(*user, *offer)
	if err != nil {
		return errors.New("create offer error\n" + err.Error())
	}
	return nil
}
