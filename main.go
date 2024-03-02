package main

import (
	"fmt"
	"gameapp/adapter/mysql"
	"gameapp/config"
	"gameapp/delivery/httpserver"
	"gameapp/repository/migrator"
	"gameapp/repository/mysql/mysqluser"
	"gameapp/service/authservice"
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
	fmt.Println(cfg)

	authSvc, userSvc, userValidator := setupServices(cfg)
	server := httpserver.New(cfg, authSvc, userSvc, userValidator)

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

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)
	MysqlAdapter := mysql.New(cfg.Mysql)
	MysqlRepo := mysqluser.New(MysqlAdapter)
	uV := uservalidator.New(&MysqlRepo)

	userSvc := userservice.New(authSvc, &MysqlRepo)
	return authSvc, userSvc, uV
}
