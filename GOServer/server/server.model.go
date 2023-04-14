package serverModel

import (
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
func (s *Server) StartProxy() {
	s.echoServer.Use(middleware.Recover())
	s.echoServer.Use(middleware.CORS())
	s.echoServer.Use(middleware.Gzip())
	s.echoServer.Use(middleware.RequestID())
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

func (s *Server) AddAuth(groupPath string, method string, methodPath string, f echo.HandlerFunc, conf echo.MiddlewareFunc) {
	route := s.echoServer.Group(groupPath)
	route.Use(conf)

	switch method {
	case "POST":
		route.POST(methodPath, echo.HandlerFunc(f))
	case "GET":
		route.GET(methodPath, echo.HandlerFunc(f))
	case "UPDATE":
		route.PUT(methodPath, echo.HandlerFunc(f))
	}
}

func (s *Server) AddAdminAuth(groupPath string, method string, methodPath string, f echo.HandlerFunc, conf echo.MiddlewareFunc, check echo.MiddlewareFunc) {
	route := s.echoServer.Group(groupPath)
	route.Use(conf)
	route.Use(check)
	switch method {
	case "POST":
		route.POST(methodPath, echo.HandlerFunc(f))
	case "GET":
		route.GET(methodPath, echo.HandlerFunc(f))
	case "UPDATE":
		route.PUT(methodPath, echo.HandlerFunc(f))
	}
}
