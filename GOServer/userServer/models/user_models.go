package models

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserJson struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pass"`
	Role     string `json:"role"`
	Type     string `json:"type"`
}

func (u *UserJson) HashPassword() error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 15)
	if err != nil {
		return err
	}
	u.Password = string(password)
	return nil
}

func (u *UserJson) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))
	return err == nil
}

type UserJwt struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	UType string `json:"user_type"`
	jwt.RegisteredClaims
}
