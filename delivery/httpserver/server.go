package httpserver

import (
	"fmt"
	"gameapp/config"
	"gameapp/delivery/httpserver/orderdelayhandler"
	"gameapp/delivery/httpserver/userhandler"
	"gameapp/service/orderdelayservice"
	"gameapp/service/userservice"
	"gameapp/validator/orderdelayvalidator"
	"gameapp/validator/uservalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config            config.Config
	userHandler       userhandler.Handler
	orderDelayHandler orderdelayhandler.Handler
	Router            *echo.Echo
}

func New(config config.Config, userSvc userservice.Service, userValidator uservalidator.Validator, orderDelaySvc orderdelayservice.Service, orderDelayValidator orderdelayvalidator.Validator) Server {
	return Server{
		config:            config,
		userHandler:       userhandler.New(userSvc, userValidator),
		orderDelayHandler: orderdelayhandler.New(orderDelaySvc, orderDelayValidator),
		Router:            echo.New(),
	}
}

func (s Server) Serve() *echo.Echo {
	// Middleware
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

	s.userHandler.SetUserRoutes(s.Router)
	s.orderDelayHandler.SetOrderDelayRoutes(s.Router)
	address := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
	return s.Router
}
