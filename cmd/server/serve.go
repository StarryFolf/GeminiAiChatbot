package server

import (
	"fiber/pkg/config"
	"fiber/pkg/middleware"
	"fiber/pkg/route"
	"fiber/platform/database"
	"fiber/platform/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	appCfg := config.AppConfig()

	logger.SetUpLogger()
	logr := logger.GetLogger()

	// Connect to DB.
	if err := database.ConnectDB(); err != nil {
		logr.Panicf("failed database setup. error: %v", err)
	}

	////Drop tables.
	//if err := database.DropTables(); err != nil {
	//	logr.Panicf("failed drop tables. error: %v", err)
	//}
	//
	////Migrate DB.
	//if err := database.Migrate(); err != nil {
	//	logr.Panicf("failed migrate database. error: %v", err)
	//}
	//
	//// Seed data.
	//if err := database.SeedData(); err != nil {
	//	logr.Panicf("failed seed data. error: %v", err)
	//}

	// Define Fiber config & app.
	fiberCfg := config.FiberConfig()
	app := fiber.New(fiberCfg)

	// Attach Middlewares.
	middleware.FiberMiddleware(app)

	// Routes.
	route.ConversationRoute(app)
	route.MessageRoute(app)

	// Start http server.
	serverAddr := fmt.Sprintf("%s:%d", appCfg.Host, appCfg.Port)
	if err := app.Listen(serverAddr); err != nil {
		logr.Errorf("Oops... server is not running! error: %v", err)
	}
}
