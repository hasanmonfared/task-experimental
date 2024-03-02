package config

import (
	"gameapp/adapter/mysql"
	"gameapp/adapter/redis"
	"gameapp/service/authservice"
	"time"
)

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type HTTPServer struct {
	Port int `koanf:"port"`
}
type Config struct {
	Application Application        `koanf:"application"`
	HTTPServer  HTTPServer         `koanf:"http_server"`
	Auth        authservice.Config `koand:"auth"`
	Mysql       mysql.Config       `koanf:"mysql"`
	Redis       redis.Config       `koanf:"redis"`
}
