package httpserver

import (
	"fmt"
	"gameapp/config"
	"gameapp/delivery/httpserver/delayreporthandler"
	"gameapp/delivery/httpserver/userhandler"
	"gameapp/service/delayreportservice"
	"gameapp/service/userservice"
	"gameapp/validator/delayreportvalidator"
	"gameapp/validator/uservalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config             config.Config
	userHandler        userhandler.Handler
	delayReportHandler delayreporthandler.Handler
	Router             *echo.Echo
}

func New(config config.Config, userSvc userservice.Service, userValidator uservalidator.Validator, delayReportSvc delayreportservice.Service, delayReportValidator delayreportvalidator.Validator) Server {
	return Server{
		config:             config,
		userHandler:        userhandler.New(userSvc, userValidator),
		delayReportHandler: delayreporthandler.New(delayReportSvc, delayReportValidator),
		Router:             echo.New(),
	}
}

func (s Server) Serve() *echo.Echo {
	// Middleware
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

	s.userHandler.SetUserRoutes(s.Router)
	s.delayReportHandler.SetDelayReportRoutes(s.Router)
	address := fmt.Sprintf(":%d", s.config.HTTPServer.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
	return s.Router
}
