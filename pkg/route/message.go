package route

import (
	"fiber/app/controller"
	"github.com/gofiber/fiber/v2"
)

func MessageRoute(app *fiber.App) {
	artworkGroup := app.Group("/messages")
	{
		artworkGroup.Post("/", controller.CreateMessage)
	}
}
