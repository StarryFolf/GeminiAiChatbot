package main

import (
	"fiber/cmd/server"
	"fiber/pkg/config"
)

func main() {
	config.LoadAllConfigs(".env")
	server.Serve()
}
