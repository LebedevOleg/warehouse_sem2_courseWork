package transport

import (
	"fmt"
	"net/http"
	"practice2sem/itemsServer/database"
	"practice2sem/itemsServer/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateItem(ctx echo.Context) error {
	item := new(models.ItemJson)
	err := ctx.Bind(item)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db, err := database.GetPostgresql()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Ошибка доступа к БД"+err.Error())
	}
	err = db.CreateItem(*item)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Ошибка создания объекта"+err.Error())
	}
	return ctx.JSON(http.StatusCreated, item)
}

func GetItem(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Println("\n", id)
	db, err := database.GetPostgresql()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Ошибка доступа к БД"+err.Error())
	}
	itemRow := db.GetItem(id)
	var resultItem models.ItemJson
	err = itemRow.Scan(
		&resultItem.Id, &resultItem.Name,
		&resultItem.Descriptions, &resultItem.Category_id,
		&resultItem.Category_name, &resultItem.Price_for_unit,
		&resultItem.Dimension)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "ошибка получения данных "+err.Error())
	}
	return ctx.JSON(http.StatusOK, resultItem)
}

func GetAllItems(ctx echo.Context) error {
	db, err := database.GetPostgresql()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Ошибка доступа к БД"+err.Error())
	}
	rows, err := db.GetAllItems()
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Ошибка получения данных "+err.Error())
	}
	itemsArr := make([]models.ItemJson, 0, 256)
	i := 0
	for rows.Next() {
		itemsArr = append(itemsArr, models.ItemJson{})
		err = rows.Scan(
			&itemsArr[i].Id, &itemsArr[i].Name,
			&itemsArr[i].Descriptions, &itemsArr[i].Category_id,
			&itemsArr[i].Category_name, &itemsArr[i].Price_for_unit,
			&itemsArr[i].Dimension)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "ошибка получения данных "+err.Error())
		}
		i++
	}
	return ctx.JSON(http.StatusOK, echo.Map{"allItems": itemsArr})
}

func UpdateItem(ctx echo.Context) error {}
