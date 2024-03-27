package config

import (
	"os"
	"strconv"
	"time"
)

type App struct {
	Host        string
	Port        int
	Debug       bool
	ReadTimeout time.Duration
}

var app = &App{}

func AppConfig() *App {
	return app
}

func LoadApp() {
	app.Host = os.Getenv("APP_HOST")
	app.Port, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	app.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))
	timeout, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	app.ReadTimeout = time.Duration(timeout) * time.Second
}
