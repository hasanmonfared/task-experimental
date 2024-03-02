package config

import (
	"gameapp/adapter/mysql"
	"gameapp/adapter/redis"
	"time"
)

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type HTTPServer struct {
	Port int `koanf:"port"`
}
type Config struct {
	Application Application  `koanf:"application"`
	HTTPServer  HTTPServer   `koanf:"http_server"`
	Mysql       mysql.Config `koanf:"mysql"`
	Redis       redis.Config `koanf:"redis"`
}
