package main

import (
	"fmt"
	"gameapp/adapter/estimate"
	"gameapp/adapter/mysql"
	"gameapp/config"
	"gameapp/delivery/httpserver"
	"gameapp/repository/migrator"
	"gameapp/repository/mysql/mysqlorderdelay"
	"gameapp/repository/mysql/mysqltrip"
	"gameapp/repository/mysql/mysqluser"
	"gameapp/service/orderdelayservice"
	"gameapp/service/tripservice"
	"gameapp/service/userservice"
	"gameapp/validator/uservalidator"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.Load("./config.yml")
	mgr := migrator.New(cfg.Mysql)
	mgr.Up()

	userSvc, userValidator, orderDelaySvc := setupServices(cfg)
	server := httpserver.New(cfg, userSvc, userValidator, orderDelaySvc)

	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, cfg.Application.GracefulShutdownTimeout)
	defer cancel()
	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)

	}
	time.Sleep(cfg.Application.GracefulShutdownTimeout)
	<-ctxWithTimeout.Done()
}

func setupServices(cfg config.Config) (userservice.Service, uservalidator.Validator, orderdelayservice.Service) {
	mysqlAdapter := mysql.New(cfg.Mysql)

	mysqlOrderDelay := mysqlorderdelay.New(mysqlAdapter)
	mysqlTrip := mysqltrip.New(mysqlAdapter)
	tripSvc := tripservice.New(mysqlTrip)
	estimateClient := estimate.New("")
	orderDelaySvc := orderdelayservice.New(mysqlOrderDelay, tripSvc, estimateClient)

	mysqlUser := mysqluser.New(mysqlAdapter)
	uV := uservalidator.New(&mysqlUser)
	userSvc := userservice.New(&mysqlUser)
	return userSvc, uV, orderDelaySvc
}
