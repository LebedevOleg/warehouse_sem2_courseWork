package main

import serverModel "practice2sem/server"

func main() {
	server := serverModel.NewServer(":8003")

	server.Start()
}
