package transport

import (
	"net/http"
	"practice2sem/userServer/models"
	"practice2sem/userServer/services"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Registr(ctx echo.Context) error {
	user := new(models.UserJson)
	err := ctx.Bind(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = services.Register(user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, user)
}

/* err = user.HashPassword()
if err != nil {
	return ctx.String(http.StatusBadRequest, "Ошибка при хэшировании пароля")
}
db, err := database.GetPostgresql()
if err != nil {
	return ctx.String(http.StatusBadRequest, "Ошибка доступа к базе данных")
}
err = db.CreateUser(*user) */

func Login(ctx echo.Context) error {
	user := new(models.UserJson)
	err := ctx.Bind(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	t, err := services.Login(user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, echo.Map{
		"token": t,
	})
}

/* db, err := database.GetPostgresql()
if err != nil {
	return ctx.String(http.StatusBadRequest, "Ошибка доступа к базе данных")
}
row := db.GetUser(*user)
var password string
var email string
var uType string
err = row.Scan(&email, &password, &uType)
if err != nil {
	return ctx.String(http.StatusBadRequest, "ошибка получения пароля "+err.Error())
}
if !user.CheckPasswordHash(password) {
	return ctx.String(http.StatusBadRequest, "Неверный пароль "+password)
}
//! Сделать создание ключа доступа
//* jwt token
log.Println(email, "\n", uType)
claims := &models.UserJwt{
	email,
	uType,
	jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 4))},
}
userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
t, err := userToken.SignedString([]byte("secret"))
if err != nil {
	return ctx.String(http.StatusBadRequest, "Ошибка создания токена \n"+err.Error())
} */

// test auth
func CheckAuth(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.UserJwt)
	email := claims.Email
	return ctx.String(http.StatusOK, "welcom "+email+"\n"+claims.UType+"\n4"+claims.ExpiresAt.String())
}

// Проверка необходимого статуса.
// status - имя статуса которому должен соответствовать пользователь
func CheckStatus(status string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			user := ctx.Get("user").(*jwt.Token)
			claims := user.Claims.(*models.UserJwt)
			if claims.UType == status {
				return next(ctx)
			}
			return ctx.String(http.StatusBadRequest, "Недостаточный статус. Ваш: \n"+claims.UType+"\nНужен: "+status)
		}
	}
}

// Создает конфиг для того чтобы проанализировать пользователя.
func CreateConfig(key string) echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.UserJwt)
		},
		SigningKey: []byte(key),
	}
	return echojwt.WithConfig(config)
}
