package main

import (
	"practice2sem/itemsServer/transport"
	serverModel "practice2sem/server"
)

func main() {
	server := serverModel.NewServer(":8002")

	server.AddRoute("/check/additem", "POST", "", transport.CreateItem)
	server.AddRoute("/check/updateitem", "POST", "", transport.UpdateItem)
	server.AddRoute("/api/get/:id", "GET", "", transport.GetItem)
	server.AddRoute("/api/getallitems", "GET", "", transport.GetAllItems)
	server.AddRoute("/api/getallcategories", "GET", "", transport.GetItemCategories)
	server.AddRoute("/check/getallstocks", "GET", "", transport.GetAllStocks)
	server.AddRoute("/check/createstock", "POST", "", transport.CreateStock)
	server.Start()
}
