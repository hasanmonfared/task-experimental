package main

import (
	"fmt"
	"gameapp/adapter/estimate"
	"gameapp/adapter/mysql"
	"gameapp/config"
	"gameapp/delivery/httpserver"
	"gameapp/repository/migrator"
	"gameapp/repository/mysql/mysqldelayreport"
	"gameapp/repository/mysql/mysqlorder"
	"gameapp/repository/mysql/mysqltrip"
	"gameapp/repository/mysql/mysqluser"
	"gameapp/service/delayreportservice"
	"gameapp/service/tripservice"
	"gameapp/service/userservice"
	"gameapp/validator/delayreportvalidator"
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

	userSvc, userValidator, delayReportSvc, delayReportValidator := setupServices(cfg)
	server := httpserver.New(cfg, userSvc, userValidator, delayReportSvc, delayReportValidator)

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

func setupServices(cfg config.Config) (userservice.Service, uservalidator.Validator, delayreportservice.Service, delayreportvalidator.Validator) {
	// MYSQL
	mysqlAdapter := mysql.New(cfg.Mysql)
	// Order Delay
	mysqldelayReport := mysqldelayreport.New(mysqlAdapter)
	mysqlTrip := mysqltrip.New(mysqlAdapter)
	tripSvc := tripservice.New(mysqlTrip)
	estimateClient := estimate.New(config.UrlForEstimateClient)
	delayReportSvc := delayreportservice.New(mysqldelayReport, tripSvc, estimateClient)
	mysqlOrder := mysqlorder.New(mysqlAdapter)
	orderV := delayreportvalidator.New(mysqlOrder)
	// User
	mysqlUser := mysqluser.New(mysqlAdapter)
	uV := uservalidator.New(&mysqlUser)
	userSvc := userservice.New(&mysqlUser)
	return userSvc, uV, delayReportSvc, orderV
}
