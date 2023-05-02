package transport

import (
	"net/http"
	"practice2sem/transactionsServer/services"

	"github.com/labstack/echo/v4"
)

func AddDelivery(ctx echo.Context) error {
	return nil
}

func GetProviders(ctx echo.Context) error {
	providers, err := services.GetAllProviders()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"providers": providers})
}
