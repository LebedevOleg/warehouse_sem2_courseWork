package transport

import (
	"net/http"
	"practice2sem/transactionsServer/models"
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

func GetStorages(ctx echo.Context) error {
	storages, err := services.GetAllStorages()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"storages": storages})
}

func CreateDelivery(ctx echo.Context) error {
	deliveryData := new(models.DeliveryRequest)
	err := ctx.Bind(deliveryData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	file, err := services.CreateNewDelivery(*deliveryData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.File(file)
}
