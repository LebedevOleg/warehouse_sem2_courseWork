package services

import (
	"errors"
	"practice2sem/userServer/database"
	"practice2sem/userServer/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Register(user *models.UserJson) error {
	err := user.HashPassword()
	if err != nil {
		return errors.New("hash password error\n" + err.Error())
	}
	db, err := database.GetPostgresql()
	if err != nil {
		return errors.New("get postgresql error\n" + err.Error())
	}
	err = db.CreateUser(*user)
	if err != nil {
		return errors.New("create user error\n" + err.Error())
	}
	return nil
}

func Login(user *models.UserJson) (string, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return "", errors.New("get postgresql error\n" + err.Error())
	}
	row := db.GetUser(*user)
	var password string
	var email string
	var uType string
	err = row.Scan(&email, &password, &uType)
	if err != nil {
		return "", errors.New("get user error\n" + err.Error())
	}
	if !user.CheckPasswordHash(password) {
		return "", errors.New("password unmatch\n")
	}
	claims := &models.UserJwt{
		email,
		uType,
		jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 4))},
	}
	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := userToken.SignedString([]byte("secret"))
	if err != nil {
		return "", errors.New("sign token error\n" + err.Error())
	}
	return t, nil
}
