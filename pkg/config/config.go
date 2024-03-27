package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func LoadAllConfigs(envFile string) {
	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatalf("Can't load env file: %v", err)
	}

	LoadApp()
	LoadDbCfg()
	loadGeminiCfg()
}

func FiberConfig() fiber.Config {
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(AppConfig().ReadTimeout),
	}
}
