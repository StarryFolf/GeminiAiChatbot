package route

import (
	"fiber/app/controller"
	"github.com/gofiber/fiber/v2"
)

func ConversationRoute(app *fiber.App) {
	artworkGroup := app.Group("/conversations")
	{
		artworkGroup.Post("/", controller.CreateConversation)
		artworkGroup.Get("/", controller.GetAllConversations)
		artworkGroupId := artworkGroup.Group("/:conversationId")
		{
			artworkGroupId.Get("/", controller.GetConversationById)
			artworkGroupId.Get("/messages", controller.GetMessagesOfConversation)
		}
	}
}
