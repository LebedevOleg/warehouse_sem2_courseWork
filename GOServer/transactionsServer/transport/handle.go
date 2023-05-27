package transport

import (
	"fmt"
	"net/http"
	"path/filepath"
	"practice2sem/transactionsServer/models"
	"practice2sem/transactionsServer/services"
	"strconv"

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
	data, err := services.CreateNewDelivery(*deliveryData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, echo.Map{"data": data})
}

func GetFileTemp(ctx echo.Context) error {
	return ctx.File(filepath.Join(".", "templates", "Delivery_temp.docx"))
}

func CreateProvider(ctx echo.Context) error {
	providerData := new(models.Provider)
	err := ctx.Bind(providerData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = services.CreateNewProvider(*providerData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func GetAllOrders(ctx echo.Context) error {
	orders, err := services.GetAllOrders()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"orders": orders})
}

func GetUsersOrders(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Request().Header.Get("id"))
	if err != nil {
		return err
	}
	fmt.Println(id)
	orders, err := services.GetUsersOrders(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"orders": orders})
}

func UpdateOrderStatus(ctx echo.Context) error {
	order := new(models.OrderJson)
	err := ctx.Bind(order)
	if err != nil {
		return err
	}
	err = services.UpdateOrderStatus(order.Id, order.Status)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func GetOrderDetails(ctx echo.Context) error {
	orderId := new(models.OrderJson)
	err := ctx.Bind(orderId)
	if err != nil {
		return err
	}
	items, err := services.GetOrderDetails(orderId.Id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, echo.Map{"items": items})
}

func UpdateOrder(ctx echo.Context) error {
	order := new(models.OrderJson)
	err := ctx.Bind(order)
	if err != nil {
		return err
	}
	err = services.UpdateOrder(*order)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func DeleteOrder(ctx echo.Context) error {
	orderId := new(models.OrderJson)
	err := ctx.Bind(orderId)
	if err != nil {
		return err
	}
	err = services.DeleteOrder(orderId.Id)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}
