package main

import (
	"net/http/httputil"
	"net/url"
	serverModel "practice2sem/server"

	"github.com/labstack/echo/v4"
)

func main() {
	proxyServer := serverModel.NewServer(":8000")

	userProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8001",
		Path:   "/api",
	})
	itemAuthProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8001",
		Path:   "/check",
	})
	itemNotauthProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8002",
		Path:   "/api",
	})

	proxyServer.AddRoute("", "POST", "/login", echo.WrapHandler(userProxy))
	proxyServer.AddRoute("", "POST", "/registr", echo.WrapHandler(userProxy))
	proxyServer.AddRoute("", "POST", "/additem", echo.WrapHandler(itemAuthProxy))
	proxyServer.AddRoute("", "GET", "/get/:id", echo.WrapHandler(itemNotauthProxy))
	proxyServer.AddRoute("", "GET", "/getall", echo.WrapHandler(itemNotauthProxy))
	proxyServer.StartProxy()
}
