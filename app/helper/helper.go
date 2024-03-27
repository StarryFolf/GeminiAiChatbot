package helper

import (
	"fiber/app/models"
	"github.com/gofiber/fiber/v2"
)

func SystemError(ctx *fiber.Ctx) error {
	response := models.BaseResponseWithErrors{
		ResultMessage: "System error",
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(response)
}

func ResponseWithError(ctx *fiber.Ctx, statusCode int, errMsg string) error {
	response := models.BaseResponseWithErrors{
		ResultMessage: errMsg,
	}

	return ctx.Status(statusCode).JSON(response)
}
