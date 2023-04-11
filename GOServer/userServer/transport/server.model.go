package transport

import (
	"practice2sem/userServer/models"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(address string) Server {
	server := Server{address: address}
	server.echoServer = echo.New()
	return server
}

type Server struct {
	address    string
	echoServer *echo.Echo
}

func (s *Server) Start() {
	s.echoServer.Use(middleware.Logger())
	s.echoServer.Logger.Fatal(s.echoServer.Start(s.address))
}

//type echoRoute func(echo.Context) error

func (s *Server) AddRoute(groupPath string, method string, methodPath string, f echo.HandlerFunc) {
	route := s.echoServer.Group(groupPath)
	switch method {
	case "POST":
		route.POST(methodPath, echo.HandlerFunc(f))
	case "GET":
		route.GET(methodPath, echo.HandlerFunc(f))
	case "UPDATE":
		route.PUT(methodPath, echo.HandlerFunc(f))
	}
}

func (s *Server) AddAuth(groupPath string, method string, methodPath string, f echo.HandlerFunc) {
	route := s.echoServer.Group(groupPath)
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.UserJwt)
		},
		SigningKey: []byte("secret"),
	}
	route.Use(echojwt.WithConfig(config))

	switch method {
	case "POST":
		route.POST(methodPath, echo.HandlerFunc(f))
	case "GET":
		route.GET(methodPath, echo.HandlerFunc(f))
	case "UPDATE":
		route.PUT(methodPath, echo.HandlerFunc(f))
	}
}
