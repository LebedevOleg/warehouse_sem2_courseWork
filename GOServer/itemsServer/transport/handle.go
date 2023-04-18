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
