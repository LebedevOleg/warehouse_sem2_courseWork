package main

import (
	serverModel "practice2sem/server"
	"practice2sem/transactionsServer/transport"
)

func main() {
	server := serverModel.NewServer(":8003")

	server.AddRoute("/check/allproviders", "GET", "", transport.GetProviders)
	server.AddRoute("/check/allstorages", "GET", "", transport.GetStorages)
	//server.AddRoute()
	server.Start()
}
