package transport

import (
	"net/http"
	"practice2sem/itemsServer/models"
	"practice2sem/itemsServer/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateItem(ctx echo.Context) error {
	item := new(models.ItemJson)
	err := ctx.Bind(item)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = services.CreateItem(item)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, "Created successfully")
}

func GetItem(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	item, err := services.GetItem(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, *item)
}

func GetAllItems(ctx echo.Context) error {
	items, err := services.GetAllItems()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, echo.Map{"allItems": items})
}

func GetItemCategories(ctx echo.Context) error {
	categories, err := services.GetAllCategories()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, echo.Map{"categories": categories})
}

func UpdateItem(ctx echo.Context) error {
	item := new(models.ItemJson)
	err := ctx.Bind(item)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = services.UpdateItem(item)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK,
		echo.Map{"message": "Updated successfully", "item": item})
}

func GetAllStocks(ctx echo.Context) error {
	stocks, err := services.GetAllStocks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, echo.Map{"stocks": stocks})
}

func CreateStock(ctx echo.Context) error {
	stock := new(models.StockJson)
	err := ctx.Bind(stock)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = services.CreateStock(stock)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, "Created successfully")
}

func CreatePurchase(ctx echo.Context) error {
	purchase := new(models.PurchaseJson)
	err := ctx.Bind(purchase)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = services.CreatePurchase(purchase)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, "Created successfully")
}
