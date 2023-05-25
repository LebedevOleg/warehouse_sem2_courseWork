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
	config := transport.CreateConfig("secret")

	server.AddRoute("/api", "POST", "/registr", transport.Registr)
	server.AddRoute("/api", "POST", "/login", transport.Login)
	server.AddAdminAuth("/api", "GET", "/getallusers", transport.GetAllUsers, config, transport.CheckStatus("Администратор"))
	server.AddAdminAuth("/api", "GET", "/getuserinfo", transport.GetUserInfo, config, transport.CheckLogin())
	server.AddAdminAuth("/api", "GET", "/updateuser", transport.UpdateUserInfo, config, transport.CheckLogin())
	server.AddAdminAuth("/api", "POST", "/createoffer", transport.CreateOffer, config, transport.CheckLogin())

	//--------------------------------routes items-auth routes------------------------------------------
	itemProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8002",
	})
	server.AddAdminAuth("/check", "GET", "/getallstocks",
		echo.WrapHandler(itemProxy), config, transport.CheckStatus("Администратор"))
	server.AddAdminAuth("/check", "POST", "/additem",
		echo.WrapHandler(itemProxy), config, transport.CheckStatus("Администратор"))
	server.AddAdminAuth("/check", "POST", "/updateitem",
		echo.WrapHandler(itemProxy), config, transport.CheckStatus("Администратор"))
	server.AddAdminAuth("/check", "GET", "/createstock",
		echo.WrapHandler(itemProxy), config, transport.CheckStatus("Администратор"))

	//--------------------------------routes transactions-auth routes------------------------------------------
	transactionsProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8003",
	})
	server.AddAdminAuth("/check", "GET", "/allproviders",
		echo.WrapHandler(transactionsProxy), config, transport.CheckStatus("Администратор"))
	server.AddRoute("/check", "GET", "/allstorages", echo.WrapHandler(transactionsProxy))
	server.AddAdminAuth("/check", "POST", "/addtransaction",
		echo.WrapHandler(transactionsProxy), config, transport.CheckStatus("Администратор"))
	server.AddAdminAuth("/check", "POST", "/createprovider",
		echo.WrapHandler(transactionsProxy), config, transport.CheckStatus("Администратор"))
	server.AddAdminAuth("/check", "GET", "/getfile",
		echo.WrapHandler(transactionsProxy), config, transport.CheckStatus("Администратор"))

	server.Start()

}
