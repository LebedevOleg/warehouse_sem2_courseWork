package main

import (
	"practice2sem/itemsServer/transport"
	serverModel "practice2sem/server"
)

func main() {
	server := serverModel.NewServer(":8002")

	server.AddRoute("/check/additem", "POST", "", transport.CreateItem)
	server.AddRoute("/api/get/:id", "GET", "", transport.GetItem)

	server.Start()
}
