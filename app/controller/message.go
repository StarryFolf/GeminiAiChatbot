package controller

import (
	"fiber/app/helper"
	"fiber/app/models"
	"fiber/app/models/entities"
	"fiber/app/models/requests"
	"fiber/pkg/config"
	"fiber/pkg/utils"
	"fiber/platform/database"
	"fiber/platform/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/generative-ai-go/genai"
	"github.com/google/uuid"
	"time"
)

func CreateMessage(ctx *fiber.Ctx) error {
	db := database.GetDB()
	logger.SetUpLogger()
	logr := logger.GetLogger()
	geminiClient, context := config.GeminiModel()
	defer geminiClient.Close()

	geminiModel := geminiClient.GenerativeModel("gemini-pro")
	geminiModel.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
	}

	request := &requests.CreateMessageRequest{}
	err := ctx.BodyParser(request)

	if err != nil {
		logr.Errorf("Failed to parse body: %v", err)
		return helper.ResponseWithError(ctx, fiber.StatusBadRequest, "Failed to parse the request!")
	}

	conversationId, err := uuid.Parse(request.Data.ConversationId)
	if err != nil {
		logr.Errorf("Failed to parse conversation Id: %v", err)
		return helper.ResponseWithError(ctx, fiber.StatusBadRequest, "Failed to parse the request!")
	}

	validate := utils.NewValidator()

	if err = validate.Struct(request); err != nil {
		logr.Errorf("Failed to validate fields: %v", err)
		return helper.ResponseWithError(ctx, fiber.StatusBadRequest, "Failed to validate the request!")
	}

	var messages []*entities.Message
	condition := "conversation_id = ?"
	db.Model(messages).Where(condition, conversationId).Find(&messages).Order("sent_time ASC")

	var chatHistory []*genai.Content

	var chatSession = geminiModel.StartChat()

	for _, message := range messages {
		chatContent := &genai.Content{
			Parts: []genai.Part{
				genai.Text(message.Content),
			},
			Role: message.Sender,
		}

		chatHistory = append(chatHistory, chatContent)
	}

	chatSession.History = chatHistory

	chatResponse, err := chatSession.SendMessage(*context, genai.Text(request.Data.Content))

	if err != nil {
		logr.Errorf("Something went wrong with the process: %v", err)
		return helper.SystemError(ctx)
	}

	messageId := uuid.New()
	message := &entities.Message{
		Id:             messageId,
		Content:        request.Data.Content,
		Sender:         "user",
		SentTime:       time.Now(),
		ConversationId: &conversationId,
	}

	err = db.Create(&message).Error
	if err != nil {
		logr.Errorf("Failed to create message. Error: %v", err)
		return helper.SystemError(ctx)
	}

	responseMessageId := uuid.New()
	responseMessage := &entities.Message{
		Id:             responseMessageId,
		Content:        fmt.Sprintf("%v", chatResponse.Candidates[0].Content.Parts[0]),
		Sender:         "model",
		ConversationId: &conversationId,
		SentTime:       time.Now(),
	}

	err = db.Create(&responseMessage).Error
	if err != nil {
		logr.Errorf("Failed to create response message. Error: %v", err)
		return helper.SystemError(ctx)
	}

	response := models.BaseResponse{
		ResultMessage: "Success!",
		Data:          responseMessage,
	}

	logr.WithField("response", response).Info("Create message response")
	return ctx.Status(fiber.StatusOK).JSON(response)
}
