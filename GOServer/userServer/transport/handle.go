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
	t, role, err := services.Register(user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, echo.Map{
		"token": t,
		"role":  role,
	})
}

func GetAllUsers(ctx echo.Context) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, users)
}

func Login(ctx echo.Context) error {
	user := new(models.UserJson)
	err := ctx.Bind(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	t, role, err := services.Login(user)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusAccepted, echo.Map{
		"token": t,
		"role":  role,
	})
}

func GetUserInfo(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.UserJwt)
	id := claims.Id
	userInfo, err := services.GetUserInfo(id)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, userInfo)
}

func UpdateUserInfo(ctx echo.Context) error {
	upUser := new(models.UserJson)
	err := ctx.Bind(upUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = services.UpdateUserInfo(upUser)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func CreateOffer(ctx echo.Context) error {
	jwt := ctx.Get("user").(*jwt.Token)
	user := jwt.Claims.(*models.UserJwt)
	offerItems := make([]models.Item, 0, 256)
	err := ctx.Bind(&offerItems)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = services.CreateOffer(user, offerItems)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

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

func CheckLogin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			user := ctx.Get("user").(*jwt.Token)
			claims := user.Claims.(*models.UserJwt)
			if claims.Email != "" {
				return next(ctx)
			}
			return ctx.String(http.StatusBadRequest, "Необходимо зайти в профиль")
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
