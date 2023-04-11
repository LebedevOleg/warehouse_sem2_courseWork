package main

import (
	"net/http/httputil"
	"net/url"
	serverModel "practice2sem/server"
	"practice2sem/userServer/transport"

	"github.com/labstack/echo/v4"
)

func main() {
	server := serverModel.NewServer(":8001")

	server.AddRoute("/api", "POST", "/registr", transport.Registr)
	server.AddRoute("/api", "POST", "/login", transport.Login)
	itemProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8002",
	})
	config := transport.CreateConfig("secret")
	server.AddAdminAuth("/check", "POST", "/additem",
		echo.WrapHandler(itemProxy), config, transport.CheckStatus("Администратор"))

	server.Start()

	/* server := transport.NewServer(":8001")
	server.AddRoute("/api", "POST", "/registr", transport.Registr)
	server.AddRoute("/api", "POST", "/login", transport.Login)

	server.AddAuth("/api", "GET", "/check", transport.CheckAuth)

	server.Start()
	*/
}
